package pkg_encoding

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func DemoHex() {
	fmt.Println("hex包实现了16进制字符表示的编解码。")
	// demoHexDecode()
	// demoDecodeString()
	// demoHexDump()
	// demoHexDumper()
	demoHexEncodeToString()
}

func demoHexDecode() {
	fmt.Println(" Decode 将 src 解码为 DecodedLen(len(src)) 字节，返回写入 dst 的实际字节数。")
	src := []byte("48656c6c6f20476f7068657221")
	// 根据 decodedLen 初始化目标切片
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])
}

func demoDecodeString() {
	fmt.Println("DecodeString 返回由十六进制字符串 s 表示的字节.")
	const s = "48656c6c6f20476f7068657221"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)
}

func demoHexDump() {
	fmt.Println("dump 转储返回一个包含给定数据的十六进制转储的字符串.")
	fmt.Println("转储的格式与hexdump -C命令行上的输出相匹配。")
	content := []byte("Go is an open source programming language.")
	str := hex.Dump(content)
	fmt.Printf("%s", str)
}

func demoHexDumper() {
	fmt.Println("dumper 返回一个 WriteCloser，它将所有写入数据的十六进制转储写入 w")
	fmt.Println("转储的格式与hexdump -C命令行上的输出相匹配。")
	lines := []string{
		"Go is an open source programming language.",
		"\n",
		"We encourage all Go users to subscribe to golang-announce.",
	}

	stdoutDumper := hex.Dumper(os.Stdout)

	defer stdoutDumper.Close()

	for _, line := range lines {
		stdoutDumper.Write([]byte(line))
	}
}

func demoHexEncodeToString() {
	fmt.Println("EncodeToString 返回 src 的十六进制编码")
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)

}
