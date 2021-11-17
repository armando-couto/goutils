package goutils

import (
	"encoding/base64"
	"strings"
	"time"
)

/*
	EncodeStringToBase64
*/
func GeneratePasswordCurrent() string {
	return strings.Split(base64.StdEncoding.EncodeToString([]byte(time.Now().Format(LAYOUT_DDMMYYYY))), "=")[0]
}
