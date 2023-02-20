package pkg_container

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

// Less 对比大小 h[i]<h[j] 返回 true
func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	// Push和Pop使用指针接收器，因为它们修改切片的长度，
	// 不仅仅是其内容
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func DemoHeap() {
	fmt.Println("接口方法列表:Init初始化堆,Push插入元素,Pop删除并返回元素,Remove删除第i个元素,Fix修改第i个元素")
	fmt.Println("定义一个整数切片intHeap")
	h := &intHeap{2, 1, 5}
	fmt.Printf("初始h:%v\n", h)
	heap.Init(h) // init 从小到大
	fmt.Printf("heap init h:%v\n", h)
	heap.Push(h, 3)
	fmt.Printf("heap push h:%v\n", h)
	fmt.Printf("h0: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("pop:%d\n", heap.Pop(h))
	}
}
