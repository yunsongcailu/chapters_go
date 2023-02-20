package topic0_100

// twoSum 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	if len(nums) <= 2 {
		if nums[0]+nums[1] == target {
			return []int{0, 1}
		}
	} else if len(nums) > 2 {
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers 两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		next1 *ListNode
		next2 *ListNode

		returnNode *ListNode

		nextPtr *ListNode

		i int
	)
	next1 = l1.Next
	next2 = l2.Next
	returnNode = &ListNode{
		Val:  l1.Val + l2.Val,
		Next: nil,
	}
	if returnNode.Val-10 >= 0 {
		returnNode.Val -= 10
		i = 1
	}

	for next1 != nil || next2 != nil || i == 1 {
		newNext := &ListNode{}

		if next1 != nil {
			newNext.Val += next1.Val
			next1 = next1.Next
		}

		if next2 != nil {
			newNext.Val += next2.Val
			next2 = next2.Next
		}

		if i == 1 {
			newNext.Val += 1
			i = 0
		}

		if newNext.Val-10 >= 0 {
			newNext.Val -= 10
			i = 1
		}

		if returnNode.Next == nil {
			returnNode.Next = newNext
		} else {
			nextPtr.Next = newNext
		}
		nextPtr = newNext
	}

	return returnNode
}
