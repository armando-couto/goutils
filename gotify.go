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
	Title     string
	Message   string
	Priority  int
}

func (m MessageGotify) SendNotification() error {
	form := url.Values{}
	form.Add("title", m.Title)
	form.Add("message", m.Message)
	form.Add("priority", fmt.Sprintf("%d", m.Priority))

	resp, err := http.PostForm(fmt.Sprintf("%s/message?token=%s", m.ServerURL, m.Token), form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		CreateFileDay(Message{Error: "Falha ao enviar Notificação via Gotify: ", Objects: []interface{}{resp}})
		return errors.New(resp.Status)
	}
	return nil
}
