package goutils

import (
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

	resp, err := http.PostForm(fmt.Sprintf("%s/message?token=%s", m.ServerURL, m.Token), form)
	if err != nil {
		CreateFileDay(Message{Error: "Falha ao enviar Notificação via Gotify: ", Objects: []interface{}{resp}}, &MessageGotifyGlobal)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		CreateFileDay(Message{Error: "Não teve o Status OK do Gotify: ", Objects: []interface{}{resp}}, &MessageGotifyGlobal)
	}
}

var MessageGotifyGlobal MessageGotify
