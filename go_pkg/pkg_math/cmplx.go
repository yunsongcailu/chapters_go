package pkg_math

import (
	"fmt"
	"math"
	"math/cmplx"
)

func DemoCmplx() {
	fmt.Println("复数示例")
	euler()
}

// euler 欧拉公式
func euler() {
	fmt.Println("演示复数取模/绝对值")
	// 定义一个复数 go语言中, 4i表示虚数  3+4i表示复数的实部和虚部
	c := 3 + 4i
	// 对虚数取模 或者说是 绝对值
	cAbs := cmplx.Abs(c)
	fmt.Printf("虚数c:%v,模:%v\n", c, cAbs)
	fmt.Printf("go语言验证欧拉公式: e的iπ次方 + 1 = 0\n")
	// e:math.E,π:math.PI, i: 1i(表示i) 因为直接用i 默认解析为变量i cmplx.Pow(a,n) a的n次方
	res := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Printf("使用Pow方法求:\ne的iπ次方 + 1 =%.3f\n", res)
	// Exp 求底数为e的 n 次方
	res = cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("使用Exp方法求:\ne的iπ次方 + 1 =%.3f\n", res)
}
