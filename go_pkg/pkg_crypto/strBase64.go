package pkg_crypto

import "encoding/base64"

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
