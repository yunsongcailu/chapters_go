package pkg_encoding

import (
	"encoding/base64"
	"fmt"
	"os"
)

func DemoBase64() {
	fmt.Println("\nbase64实现了RFC 4648规定的base64编码。")
	fmt.Println("base64 EncodeToString:")
	data := []byte("any + old & data")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)

	fmt.Println("base64 DecodeString:")

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("str:%s\n,base64 encoding:%q\n", str, data)
	fmt.Println("os.Stdout标准输出 newEncoder")
	input := []byte("foo\x00bar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}
