package pkg_os

import (
	"bufio"
	"fmt"
	"os"
)

func DemoOs() {
	fmt.Println("os包提供了操作系统函数的不依赖平台的接口")
	data, err := os.ReadFile("test.csv")
	if err != nil {
		fmt.Printf("os.ReadFile error:%v\n", err)
		return
	}
	fmt.Printf("file data:%s\n", data)
	publicPem := "public.pem"
	publicKey, err := readFileByLine(publicPem)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("public key:\n%s", publicKey)
}

// readFileByLine 按行读取文件
func readFileByLine(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var res string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = res + "\n" + scanner.Text()
	}
	return res, nil
}
