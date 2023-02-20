package pkg_crypto

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func DemoRand() {
	fmt.Println("crypto/rand 包实现了一个密码安全的伪随机数生成器。 不是math包的rand")
	fmt.Println("本例从 rand.Reader 中读取10个密码安全的伪随机数，并将它们写入字节片。")
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(b)
	// 切片现在应该包含随机字节而不是仅包含零。
	fmt.Println(bytes.Equal(b, make([]byte, c)))
}
