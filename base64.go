package goutils

import "encoding/base64"

func EncodeStringToBase64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func DecodeBase64ToString(value string) string {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		CreateFileDay("error:" + err.Error())
		return "ERRONOJSON"
	}
	return string(data)
}