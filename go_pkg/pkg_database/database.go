package pkg_database

import "fmt"

func DemoDatabase() {
	fmt.Println("database包,sql包提供了保证SQL或类SQL数据库的泛用接口。\n\n使用sql包时必须注入（至少）一个数据库驱动")
	fmt.Println("驱动列表:https://golang.org/s/sqldrivers")
	fmt.Println("使用ORM这里不多介绍")
}
