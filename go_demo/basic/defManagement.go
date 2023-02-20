package basic

import "fmt"

var Students []Student

func DefManagement() {
	menu()
	input := ""
	for input != "6" {
		fmt.Println("请选择菜单:")
		_, _ = fmt.Scanln(&input)
		operator(input)
	}
}

// menu 菜单
func menu() {
	fmt.Printf("%s\n", MenuInfo)
}

// 功能选择
func operator(id string) {
	switch id {
	case "1":
		create()
		fmt.Println(Students)
	case "2":
		fmt.Println("2del")

	case "3":
		fmt.Println("3edit")

	case "4":
		stu := first()
		fmt.Println(stu)
	case "5":
		fmt.Println("5list")

	case "6":
		fmt.Println("6quit")

	default:
		fmt.Println("输入错误")
	}
	return
}

// inputInfo 输入信息
func inputInfo() (num, name string, age int) {
	fmt.Println("请输入学号:")
	_, _ = fmt.Scanln(&num)
	fmt.Println("请输入姓名:")
	_, _ = fmt.Scanln(&name)
	fmt.Println("请输入年龄:")
	_, _ = fmt.Scanln(&age)
	return
}

// create 添加
func create() {
	num, name, age := inputInfo()
	stu := Student{
		Num:  num,
		Name: name,
		Age:  age,
	}
	Students = append(Students, stu)
	return
}

// first 查找
func first() *Student {
	selected := 0
	var num, name string
	fmt.Println("学号查询[1],姓名查询[2]")
	_, _ = fmt.Scanln(&selected)
	if selected == 1 {
		fmt.Println("请输入学号:")
		_, _ = fmt.Scanln(&num)
	} else if selected == 2 {
		fmt.Println("请输入姓名:")
		_, _ = fmt.Scanln(&name)
	} else {
		fmt.Println("输入错误")
	}
	for _, student := range Students {
		if student.Num == num || student.Name == name {
			return &student
		}
	}
	return nil
}
