package pkg_crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// EncryptAES AES加密
// 创建一个cipher.Block接口。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256。
func EncryptAES(src []byte, key []byte) (string, error) {
	fmt.Printf("key len:%d\n", len(key))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	paddedSrc := aesPadding(src, block.BlockSize())
	fmt.Printf("补位后src长度:%d\n", len(paddedSrc))
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	// 初始化新密码
	encryptedSrc := make([]byte, len(paddedSrc))
	// 执行加密
	blockMode.CryptBlocks(encryptedSrc, paddedSrc)
	return base64.StdEncoding.EncodeToString(encryptedSrc), nil
}

// DecryptAES AES解密
func DecryptAES(data string, key []byte) ([]byte, error) {
	src, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	// 初始化解密密码
	decryptedSrc := make([]byte, len(src))
	// 执行解密
	blockMode.CryptBlocks(decryptedSrc, src)
	// 去除填充
	decryptedSrc = aesUnPadding(decryptedSrc)
	return decryptedSrc, nil
}

// aesPadding 填充(补位) 填充padNum个padNum值 方便去除填充
func aesPadding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	fmt.Printf("src:%s,补位:%s\n", string(src), string(pad))
	return append(src, pad...)
}

// aesUnPadding 去除填充(补位) unPadNum个unPadNum值  unPadNum值等于填充的位数
func aesUnPadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:(n - unPadNum)]
}
