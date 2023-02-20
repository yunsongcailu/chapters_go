package pkg_crypto

import (
	"crypto/md5"
	"fmt"
	"io"
)

func DemoMD5() {
	fmt.Println("md5")
	str := "hello world"
	// []byte md5
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has) // []byte 转为16进制
	fmt.Println(md5Str)

	// string md5
	w := md5.New()
	_, _ = io.WriteString(w, str)
	fmt.Println(fmt.Sprintf("%x", w.Sum(nil)))
}
