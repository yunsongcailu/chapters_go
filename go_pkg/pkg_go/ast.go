package pkg_go

import "fmt"

func DemoAstBuild() {
	fmt.Println("ast包声明用于表示Go包语法树的类型。")
	fmt.Println("许多自动化代码生成工具都离不开语法树分析，例如goimport，gomock，wire等项目都离不开语法树分析")
	fmt.Println("Package build 收集有关 Go 包的信息。 ")
}
