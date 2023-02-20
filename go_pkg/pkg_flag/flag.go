package pkg_flag

import (
	"flag"
	"fmt"
	"time"
)

func DemoFlag() {
	fmt.Println(" Flag 包实现了命令行标志解析。")
	//demoFlagParse()
	demoFlagParse2()
}

func demoFlagParse() {
	fmt.Println("flag.String()，Bool()，Int() 定义标志")
	fmt.Println("解析命令行 -name 11,默认值 12345,例如: go run ./main.go -name 123333 -flagname 222 -species myGopher -g abc")
	var ip = flag.Int("name", 12345, "help message for name")

	var flagvar int
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")

	// 示例1：名为“species”的单个字符串标志，默认值为“gopher”。
	var species = flag.String("species", "gopher", "the species we are studying")

	// 示例2：共享变量的两个标志，因此我们可以使用速记。
	// 初始化的顺序是未定义的，因此请确保两者都使用
	// 相同的默认值。 必须使用init函数设置它们。
	var gopherType string
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")

	flag.Parse()
	fmt.Println(*ip)
	fmt.Println(flagvar)
	fmt.Println(*species)
	fmt.Println(gopherType)
}

func demoFlagParse2() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
