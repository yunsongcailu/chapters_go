package pkg_encoding

import (
	"encoding/base32"
	"fmt"
	"os"
)

func DemoBase32() {
	fmt.Println("base32 EncodeToString:")
	data := []byte("any + old & data")
	str := base32.StdEncoding.EncodeToString(data)
	fmt.Println(str)

	fmt.Println("base32 DecodeString:")

	data, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("str:%s\n,base32 encoding:%q\n", str, data)
	fmt.Println("os.Stdout标准输出 newEncoder")
	input := []byte("foo\x00bar")
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}
