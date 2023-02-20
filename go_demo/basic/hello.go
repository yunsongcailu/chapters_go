package basic

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"time"
)

// DemoHello hello world
func DemoHello() {
	fmt.Println("hello world,跟随python再来一遍")
	var a, b int = 1, 0
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("go关键字:import,package\n" +
		"const,var,map,interface,type,struct,int8/16/32/64,uint,float/32/64,bool,string,chan,goto,range,func\n" +
		"switch,select,defer,fallthrough,if,else,break,continue,return,for,go")
	fmt.Println("请输入a:")
	_, _ = fmt.Scanln(&a)
	fmt.Println("请输入b:")
	_, _ = fmt.Scanln(&b)
	fmt.Println(a + b)
}

func getName(n string) string {
	var str string
	switch n {
	case "0":
		str = "石头"
	case "1":
		str = "剪刀"
	case "2":
		str = "布"
	}
	return str
}

func getNum(min, max int) string {
	r := rand.Intn(max) + min
	return strconv.Itoa(r)
}

func winner(i string) string {
	j := getNum(0, 2)
	fmt.Printf("Human get:%s,AI get:%s\n", getName(i), getName(j))
	if (i == "0" && j == "1") || (i == "1" && j == "2") || (i == "2" && j == "0") {
		return "Winner is Human"
	} else if i == j {
		return "Round is Draw"
	} else {
		return "Winner is AI"
	}
}

func Guess() {
	var input string
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Printf("猜拳游戏:[0]石头[1]剪刀[2]布,[q]退出,请输入:")
		_, _ = fmt.Scanln(&input)
		if input == "q" {
			break
		}
		if input == "0" || input == "1" || input == "2" {
			fmt.Println(winner(input))
			continue
		}
		fmt.Println("输入有误,请重新输入")
	}
	fmt.Println("结束游戏")
	return
}

type IntSlice []int

func (sl IntSlice) Len() int {
	return len(sl)
}
func (sl IntSlice) Less(x, y int) bool {
	return sl[x] < sl[y]
}
func (sl IntSlice) Swap(x, y int) {
	sl[x], sl[y] = sl[y], sl[x]
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sl IntSlice
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	sl = append(sl, nums1...)
	sl = append(sl, nums2...)
	sort.Sort(sl)
	return median(sl)
}

func median(sl IntSlice) float64 {
	length := len(sl)
	if length < 1 {
		return 0
	}
	if length == 1 {
		return float64(sl[0])
	}
	if length == 2 {
		return (float64(sl[0]) + float64(sl[1])) / 2
	}
	i, j := 0, 0
	if length%2 == 1 {
		i = length / 2
	} else {
		i = length/2 - 1
		j = length / 2
	}
	if j == 0 {
		return float64(sl[i])
	}
	return (float64(sl[i]) + float64(sl[j])) / 2
}
