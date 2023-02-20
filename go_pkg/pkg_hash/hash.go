package pkg_hash

import (
	"crypto/md5"
	"fmt"
)

func DemoHash() {
	fmt.Println("Hash 接口能对 []byte 字节切片进行操作，如果是 string 类型，需要转为 []byte 再进行操作。")
	fmt.Println("如果是字节切片，可以使用包中的内置方法（不同的包内置方法明不同）直接进行 hash 操作。")
	// 直接进行hash操作
	b := []byte("hello world")
	fmt.Printf("str:%s,func:md5\n%x\n", string(b), md5.Sum(b))

	// 也可以使用Hash接口进行操作
	hasher := md5.New()
	hasher.Write(b)
	fmt.Printf("%x\n", hasher.Sum(nil))
	fmt.Println("MD5 返回的是 32 位字符串，而使用 Sum 返回的是一个字节切片 (如果直接使用包方法调用，返回的是字节数组)，这里需要转为 16 进制显示，才能展示 32 位字符串。")
	fmt.Println("fnv 是一种简单可靠的 hash 算法,Adler-32通过求解两个16位的数值A、B实现，并将结果连结成一个32位整数.")
	fmt.Println("crc32,crc64,crypto 包里提供了一些其他的算法也实现了 hash 接口，比如 md5，sha1，sha256等。")

}
