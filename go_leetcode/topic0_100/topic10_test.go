package topic0_100

import "testing"

func TestTopic10(t *testing.T) {
	//nums := []int{3, 2, 4}
	//target := 6
	//t.Log(twoSum(nums, target))
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val:  1,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{},
			},
		},
	}
	res := addTwoNumbers(l1, l2)
	t.Log(res)
}
