package pkg_crypto

import (
	"crypto/subtle"
	"fmt"
)

func DemoSubtle() {
	fmt.Println("crypto/subtle包常用功能:字节/切片对比,复制,获取")
	fmt.Println("比较两个byte是否相同EQ")
	byteEq()
	fmt.Println("比较两个int32Eq")
	int32Eq()
	fmt.Println("比较两个非负 int 小于等于")
	intLessOrEq()
	fmt.Println("比较两个byte切片长度和内容")
	bytesCompare()
	fmt.Println("如果v == 1,则将y的内容拷贝到x；如果v == 0，x不作修改")
	bytesCopyOrNot()
	fmt.Println("如果v == 1，返回x；如果v == 0，返回y")
	intSelect()
}

// byteEq 比较两个byte是否相同EQ
func byteEq() {
	str := "abaaad"
	for i := 0; i < len(str)-1; i++ {
		x := str[i]
		y := str[i+1]
		res := subtle.ConstantTimeByteEq(x, y)
		if res == 1 {
			fmt.Printf("i:%d,v:%v = i:%d,v:%v\n", i, str[i], i+1, str[i+1])
		} else {
			fmt.Printf("i:%d,v:%v != i:%d,v:%v\n", i, str[i], i+1, str[i+1])
		}
	}
}

// int32Eq
func int32Eq() {
	var x, y int32 = 2, 2
	res := subtle.ConstantTimeEq(x, y)
	if res == 1 {
		fmt.Printf("x:%d = y:%d\n", x, y)
	} else {
		fmt.Printf("x:%d != y:%d\n", x, y)
	}
}

// intLessOrEq int小于等于
func intLessOrEq() {
	var x, y int = 1, 2
	res := subtle.ConstantTimeLessOrEq(x, y)
	if res == 1 {
		fmt.Printf("x:%d 小于等于 y:%d\n", x, y)
	} else {
		fmt.Printf("x:%d 大于 y:%d\n", x, y)
	}
}

// bytesCompare 比较两个byte切片长度和内容
func bytesCompare() {
	x := []byte("compare")
	y := []byte("Compare")
	res := subtle.ConstantTimeCompare(x, y)
	if res == 1 {
		fmt.Printf("x:%v 长度或内容等于 y:%v\n", x, y)
	} else {
		fmt.Printf("x:%v 长度或内容不等于 y:%v\n", x, y)
	}
}

// 如果v == 1,则将y的内容拷贝到x；如果v == 0，x不作修改
func bytesCopyOrNot() {
	y := []byte("abc")
	x := make([]byte, len(y))
	v := 0
	fmt.Println("v = 0 nothing")
	subtle.ConstantTimeCopy(v, x, y)
	fmt.Printf("v:%d,x:%v,y:%v\n", v, x, y)
	fmt.Println("v = 1 copy")
	v = 1
	subtle.ConstantTimeCopy(v, x, y)
	fmt.Printf("v:%d,x:%v,y:%v\n", v, x, y)
}

// 如果v == 1，返回x；如果v == 0，返回y
func intSelect() {
	y := 11
	x := 22
	v := 0
	fmt.Println("v = 0 return y")
	result := subtle.ConstantTimeSelect(v, x, y)
	fmt.Printf("v:%d,x:%d,y:%d,result:%d\n", v, x, y, result)
	fmt.Println("v = 1 return x")
	v = 1
	result = subtle.ConstantTimeSelect(v, x, y)
	fmt.Printf("v:%d,x:%d,y:%d,result:%d\n", v, x, y, result)
}
