package tests

import (
	"github.com/armando-couto/goutils"
)

func main() {
	m := goutils.MessageGotify{
		ServerURL: "http://54.146.11.58:8485",
		Token:     "AlDU1S0-NbjrRbX",
		Priority:  10,
	}

	m.SendNotification("", "")
}
