package pkg_container

import (
	"container/ring"
	"fmt"
)

func DemoRing() {
	r := ring.New(10)
	fmt.Printf("ring new 实例化环形链表 r,r.len元素个数为:%+v,%d\n", r, r.Len())
	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}
	fmt.Printf("ring 赋值i:%+v\n", r)
	fmt.Printf("循环打印ring r.value\n")
	r.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Printf("\nr.Move(n),指针移动n(3)个位置并返回指向的节点,n>0正向,n<0反向\n")
	r = r.Move(3)
	r.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Printf("\nr.Next()指向下一个节点,r.Prev()指向上一个节点\n")
	r = r.Next()
	r.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Println()
	r = r.Prev()
	r.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Printf("演示r.Link 连接两个环形链表\n把一个环形双向链表s与环形双向链表r相链接\n并返回相连前时s.Next()的值。r不能为空。\n")
	fmt.Println("如果s和r不是同一个环形链表，则相连后，值产生一个环形链表，并返回相连前时s.Next()的值")
	fmt.Println("如果s和r是同一个环形链表，但s!=r时，相连后，产生两个环形链表，并返回相连前的s.Next()")
	r1 := ring.New(10)
	r2 := ring.New(10)
	for i := 0; i < 10; i++ {
		r1.Value = i
		r1 = r1.Next()
		r2.Value = i + 20
		r2 = r2.Next()
	}
	fmt.Println("初始化r1:")
	r1.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Printf("\n初始化r2:\n")
	r2.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Printf("\n情况一,s和r不同,r3 := r1.Link(r2),输出r3:\n")
	r3 := r1.Link(r2)
	r3.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Printf("\n分别输出r1和r2\n")
	r1.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Println()
	r2.Do(func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Println("其他情况不演示了,两个ring 是否相同结果不同")
	fmt.Println("r.Unlink(n) 从r的下一个节点开始移除 d = n%r.Len(),d个元素,d=0不产生影响")
}
