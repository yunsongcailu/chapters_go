package pkg_container

import (
	"container/list"
	"fmt"
)

func DemoList() {
	fmt.Printf("list 包实现了双向链表\n")
	l := list.New()
	fmt.Printf("创建一个新链表:%v\n", l)
	e4 := l.PushBack(4)
	fmt.Printf("在链表后边添加一个元素4:%v\n新链表e4:%v\n", l, e4)
	e1 := l.PushBack(1)
	fmt.Printf("在链表前边添加一个元素1:%v\n新链表e1:%v\n", l, e1)
	l.InsertBefore(3, e4)
	fmt.Printf("在e4链表前边添加一个元素3,l:%+v\n,e4:%+v\n", l, e4)
	l.InsertAfter(2, e1)
	fmt.Printf("在e1链表后边添加一个元素2,l:%+v\n,e1:%+v\n", l, e1)
	// 遍历l 并打印
	// l.Front() e 是链表第一个元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
