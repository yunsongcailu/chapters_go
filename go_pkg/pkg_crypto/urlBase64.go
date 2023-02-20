package pkg_crypto

import "encoding/base64"

// UrlBase64Encode URL 专用编码
func UrlBase64Encode(info []byte) string {
	encodeInfo := base64.URLEncoding.EncodeToString(info)
	return encodeInfo
}

// UrlBase64Decode URL 专用解码
func UrlBase64Decode(encodeInfo string) ([]byte, error) {
	info, err := base64.URLEncoding.DecodeString(encodeInfo)
	if err != nil {
		return nil, err
	}
	return info, nil
}
