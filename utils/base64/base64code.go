package base64

import "encoding/base64"

var customEncoding = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func EnCodeBase64(str string) string {
	return customEncoding.EncodeToString([]byte(str))
}

func DeCoding(str string) (string, error) {
	decodeString, err := customEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(decodeString), nil
}
