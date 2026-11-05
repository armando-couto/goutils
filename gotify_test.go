package goutils

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

const tokenDeTeste = "token-de-teste-nao-pode-ir-pro-log"

// logDoDia isola o diretório corrente (o CreateFileDay grava AAAAMMDD.log
// nele) e devolve uma função que lê o log do dia.
func logDoDia(t *testing.T) func() string {
	t.Helper()
	t.Chdir(t.TempDir())
	return func() string {
		b, err := os.ReadFile(time.Now().Format("20060102") + ".log")
		if err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
		return string(b)
	}
}

// gotifyForaDoAr devolve a URL de um servidor que já foi fechado: a conexão
// é recusada, como com o Gotify fora do ar.
func gotifyForaDoAr(t *testing.T) string {
	t.Helper()
	srv := httptest.NewServer(http.NotFoundHandler())
	srv.Close()
	return srv.URL
}

func semPanico(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("entrou em pânico: %v", r)
		}
	}()
	f()
}

func TestSendNotificationComGotifyForaDoArNaoEntraEmPanico(t *testing.T) {
	lerLog := logDoDia(t)
	m := MessageGotify{ServerURL: gotifyForaDoAr(t), Token: tokenDeTeste}

	semPanico(t, func() { m.SendNotification("Consulta: X", "ERROR\tFile: Consulta: X") })

	log := lerLog()
	if !strings.Contains(log, "Falha ao enviar Notificação via Gotify: Post: ") {
		t.Errorf("a falha do envio não foi para o log:\n%s", log)
	}
	if strings.Contains(log, tokenDeTeste) {
		t.Errorf("o token do Gotify foi parar no log:\n%s", log)
	}
}

// O caso dos logs da FIX-15646: a consulta falha, o CreateFileDay registra e
// tenta avisar pelo Gotify, que não responde.
func TestCreateFileDayComGotifyForaDoArNaoEntraEmPanico(t *testing.T) {
	lerLog := logDoDia(t)
	antes := MessageGotifyGlobal
	t.Cleanup(func() { MessageGotifyGlobal = antes })
	MessageGotifyGlobal = MessageGotify{ServerURL: gotifyForaDoAr(t), Token: tokenDeTeste}

	semPanico(t, func() {
		CreateFileDay(Message{File: "Consulta: X", Error: `pq: column "x" does not exist`}, &MessageGotifyGlobal)
	})

	log := lerLog()
	for _, quer := range []string{`pq: column "x" does not exist`, "Falha ao enviar Notificação via Gotify"} {
		if !strings.Contains(log, quer) {
			t.Errorf("faltou %q no log:\n%s", quer, log)
		}
	}
	if strings.Contains(log, tokenDeTeste) {
		t.Errorf("o token do Gotify foi parar no log:\n%s", log)
	}
}

func TestSendNotificationComStatusDiferenteDe200(t *testing.T) {
	lerLog := logDoDia(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	semPanico(t, func() {
		MessageGotify{ServerURL: srv.URL, Token: tokenDeTeste}.SendNotification("titulo", "mensagem")
	})

	if log := lerLog(); !strings.Contains(log, "Não teve o Status OK do Gotify: 500 Internal Server Error") {
		t.Errorf("o status do Gotify não foi para o log:\n%s", log)
	}
}

func TestSendNotificationEnviaOFormularioAoGotify(t *testing.T) {
	lerLog := logDoDia(t)
	var recebido *http.Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		recebido = r
	}))
	defer srv.Close()

	MessageGotify{ServerURL: srv.URL, Token: tokenDeTeste, Priority: 7}.SendNotification("titulo", "mensagem")

	if recebido == nil {
		t.Fatal("o Gotify não recebeu nada")
	}
	if recebido.Method != http.MethodPost || recebido.URL.Path != "/message" || recebido.URL.Query().Get("token") != tokenDeTeste {
		t.Errorf("requisição %s %s", recebido.Method, recebido.URL.Path)
	}
	for campo, quer := range map[string]string{"title": "titulo", "message": "mensagem", "priority": "7"} {
		if got := recebido.PostForm.Get(campo); got != quer {
			t.Errorf("%s = %q, queria %q", campo, got, quer)
		}
	}
	if log := lerLog(); log != "" {
		t.Errorf("envio que deu certo não deveria registrar nada:\n%s", log)
	}
}

// Gotify que aceita a conexão e não responde: o envio desiste no prazo do
// cliente, em vez de prender quem registrou o erro.
func TestSendNotificationComGotifyTravadoDesisteNoPrazo(t *testing.T) {
	lerLog := logDoDia(t)
	solta := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-solta:
		case <-r.Context().Done():
		}
	}))
	defer srv.Close()
	defer close(solta)
	prazo := clienteGotify.Timeout
	t.Cleanup(func() { clienteGotify.Timeout = prazo })
	clienteGotify.Timeout = 200 * time.Millisecond

	inicio := time.Now()
	semPanico(t, func() {
		MessageGotify{ServerURL: srv.URL, Token: tokenDeTeste}.SendNotification("titulo", "mensagem")
	})

	if d := time.Since(inicio); d > 3*time.Second {
		t.Errorf("o envio levou %v; o prazo era %v", d, clienteGotify.Timeout)
	}
	log := lerLog()
	if !strings.Contains(log, "Falha ao enviar Notificação via Gotify: Post: ") || !strings.Contains(log, "Client.Timeout exceeded") {
		t.Errorf("o estouro do prazo não foi para o log:\n%s", log)
	}
	if strings.Contains(log, tokenDeTeste) {
		t.Errorf("o token do Gotify foi parar no log:\n%s", log)
	}
}

func TestSendNotificationUsaPrazoDe5Segundos(t *testing.T) {
	if clienteGotify.Timeout != 5*time.Second {
		t.Errorf("prazo do envio ao Gotify = %v, queria 5s", clienteGotify.Timeout)
	}
}

func TestSendNotificationSemGotifyConfiguradoNaoEnvia(t *testing.T) {
	lerLog := logDoDia(t)
	semPanico(t, func() { MessageGotify{}.SendNotification("titulo", "mensagem") })
	if log := lerLog(); log != "" {
		t.Errorf("sem Gotify configurado não deveria registrar nada:\n%s", log)
	}
}
