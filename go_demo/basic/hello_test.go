package basic

import "testing"

func TestDemoHello(t *testing.T) {
	//DemoHello()
	// Guess()
	n1 := []int{1, 2}
	n2 := []int{3, 4}
	res := FindMedianSortedArrays(n1, n2)
	t.Log(res)
}
