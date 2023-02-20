package pkg_math

import (
	"fmt"
	"math"
)

func DemoSqrt() {
	fmt.Println("一些常用方法:平方根(sqrt)")
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a) + float64(b*b)))
	fmt.Printf("强制类型转换int/float64,a:%d,b:%d,平方根sqrt:%d\n", a, b, c)
	fmt.Println("四舍五入算法")
	v := 3.14159265358979323846264338
	fmt.Printf("v:%f\n", v)
	v1 := math.Round(v)
	v2 := round(v, 5)
	fmt.Printf("math.Round:%v\n 自定义round取0.0001:%v\n", v1, v2)
}

// round 四舍五入
// 返回将 val 根据指定精度 pre 保留小数点后几位（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func round(val float64, pre int) float64 {
	p := math.Pow10(pre)
	return math.Floor(val*p+0.5) / p
}
