package pkg_crypto

import "fmt"

func DemoCrypto() {
	fmt.Println("crypto常用的密码算法")
	//aesDemo()
	//DemoDsa()
	//DemoHMac()
	//DemoMD5()
	//DemoRand()
	//DemoRsa()
	//DemoSubtle()
	//DemoTls()
	DemoX509()
}

// aesDemo aes加解密
func aesDemo() {
	fmt.Printf("aes加密:key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256。\n")
	src := "abcdefg123."
	key16 := "abcdefghijklmnop"
	key24 := "abcdefghijklmnop11111111"
	key32 := "abcdefghijklmnop1111111111111111"
	src16, err := EncryptAES([]byte(src), []byte(key16))
	src24, err := EncryptAES([]byte(src), []byte(key24))
	src32, err := EncryptAES([]byte(src), []byte(key32))
	if err != nil {
		fmt.Printf("encrypt aes error:%v\n", err)
		return
	}
	fmt.Printf("src16:%s\n,src24:%v\n,src32:%v\n", string(src16), src24, src32)
	deSrc16, err := DecryptAES(src16, []byte(key16))
	deSrc24, err := DecryptAES(src24, []byte(key24))
	deSrc32, err := DecryptAES(src32, []byte(key32))
	if err != nil {
		fmt.Printf("decrypt aes error:%v\n", err)
		return
	}
	fmt.Printf("desrc16:%s\n,desrc24:%s\n,desrc32:%s\n", deSrc16, deSrc24, deSrc32)
}
