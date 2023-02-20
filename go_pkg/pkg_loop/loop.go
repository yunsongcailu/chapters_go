package pkg_loop

import (
	"fmt"
	"strconv"
)

func DemoLoop() {
	fmt.Println("循环示例")
	n := -13
	res := intToBinary(n)
	fmt.Printf("int:%d,binary:%s\n", n, res)
}

// intToBinary int转二进制
func intToBinary(n int) string {
	fmt.Println("intToBinary int转二进制,负整数以32位输出")
	fmt.Printf("n:%d,fmt二进制输出:%b,自定义方法转换返回:\n", n, n)
	var result string
	trans := func(n int) string {
		var res string
		// n = n /2 简写 n /=2
		for ; n > 0; n /= 2 {
			res = strconv.Itoa(n%2) + res
		}
		return res
	}
	if n == 0 {
		return result
	} else if n > 0 {
		return trans(n)
	} else {
		temp := -n
		tempRes := trans(temp)
		length := len([]byte(tempRes))
		// 取反操作
		buf := make([]byte, 32)
		j := 0
		for i := 0; i < 31; i++ {
			if i+length == 31 {
				buf[i] = tempRes[j]
				length -= 1
				j += 1
			} else {
				buf[i] = '0'
			}
		}
		for i := 31; i >= 0; i-- {
			if buf[i] > 0 {
				if buf[i] == '0' {
					buf[i] = '1'
				} else {
					buf[i] = '0'
				}
			} else {
				buf[i] = '0'
			}
		}
		// +1操作
		// 进位标志 flag
		flag := 0
		for i := 31; i >= 0; i-- {
			if flag == 1 && buf[i] == '1' {
				buf[i] = '1'
				flag = 1
			} else if flag == 1 && buf[i] == '0' {
				buf[i] = '1'
				flag = 0
			} else if flag == 0 && buf[i] == '0' {
				buf[i] = '0'
				flag = 0
			} else {
				buf[i] = '1'
				flag = 0
			}
		}
		result = string(buf)
	}
	return result
}
