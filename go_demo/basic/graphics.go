package basic

import "fmt"

// Square 正方形
func Square() {
	for i := 0; i < 5; i++ {
		line := ""
		for j := 0; j < 5; j++ {
			line += "  *"
		}
		fmt.Println(line)
	}
}

// Triangle1 三角形1
func Triangle1() {
	for i := 1; i <= 5; i++ {
		line := ""
		for j := 1; j <= i; j++ {
			line += "*"
		}
		fmt.Println(line)
	}
}

// Triangle2 三角形2
func Triangle2() {
	for i := 1; i <= 5; i++ {
		space, line := "", ""
		for z := 1; z <= (5 - i); z++ {
			space += " "
		}
		for j := 1; j <= i; j++ {
			line += "*"
		}
		fmt.Println(space, line)
	}
}

// Pyramid1 金字塔1
func Pyramid1() {
	for i := 1; i <= 5; i++ {
		space, line := "", ""
		for z := 1; z <= (5 - i); z++ {
			space += " "
		}
		if i == 1 {
			line = "*"
		} else {
			for j := 1; j <= 2*i-1; j++ {
				line += "*"
			}
		}
		fmt.Println(space, line)
	}
}

// MultiplicationTable 乘法表
func MultiplicationTable() {
line:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
			if i == 9 && j == 9 {
				println()
				break line
			}
		}
		println()
	}
	fmt.Println("over")
}
