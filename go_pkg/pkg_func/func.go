package pkg_func

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func DemoFunc() {
	//fmt.Println("函数示例")
	//res := apply(pow, 3, 4)
	//fmt.Printf("result:%d\n", res)
	//adder := adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i, adder(i))
	//}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// apply 计算函数
func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("演示获取函数名称:1,获取函数的指针 2,读取函数名字\n")
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("传入的函数名字为:%s,a:%d,b:%d\n", opName, a, b)
	return op(a, b)
}

// 重写pow
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// adder 累加器
func adder() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// fibonacci 斐波那契
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
