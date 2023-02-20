package pkg_crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func DemoHMac() {
	fmt.Println("hmac")
	// 待校验的mac
	messageMAC := []byte("5kScGgxUH+YqZryj+BAdaTtjoBmBn7JU3F0ljYFY5eM=")
	// new
	message := []byte("aaa")
	key := []byte("11")
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedHMac := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	// 比较 messageMac 是否是 message的 合法HMac
	checked := hmac.Equal(messageMAC, []byte(expectedHMac))
	fmt.Printf("message:%v\n,expectedMac:%v\n,messageMac:%v\n,checked:%v\n", message, expectedHMac, string(messageMAC), checked)
}
