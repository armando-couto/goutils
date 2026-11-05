package goutils

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type MessageGotify struct {
	ServerURL string
	Token     string
	Priority  int
}

func (m MessageGotify) SendNotification(title, message string) {
	if m.ServerURL == "" && m.Token == "" {
		return
	}

	if title == "" || message == "" {
		return
	}

	form := url.Values{}
	form.Add("title", title)
	form.Add("message", message)
	form.Add("priority", fmt.Sprintf("%d", m.Priority))

	// As falhas do envio vão só para o arquivo de log (Gotify nil): não há
	// como avisar pelo Gotify que o Gotify falhou.
	resp, err := http.PostForm(fmt.Sprintf("%s/message?token=%s", m.ServerURL, m.Token), form)
	if err != nil {
		// Sem resposta não há Body para fechar. O defer sobre a resposta
		// nula entrava em pânico, e com o Gotify fora do ar todo erro
		// registrado pelo CreateFileDay virava panic (500 nas APIs).
		CreateFileDay(Message{Error: "Falha ao enviar Notificação via Gotify: " + erroSemURL(err)}, nil)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		CreateFileDay(Message{Error: "Não teve o Status OK do Gotify: " + resp.Status}, nil)
	}
}

// erroSemURL devolve o erro do cliente HTTP sem a URL, que leva o token do
// Gotify na query.
func erroSemURL(err error) string {
	var ue *url.Error
	if errors.As(err, &ue) {
		return ue.Op + ": " + ue.Err.Error()
	}
	return err.Error()
}

var MessageGotifyGlobal MessageGotify
