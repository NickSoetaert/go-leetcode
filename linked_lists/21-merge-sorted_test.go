package linked_lists

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	l1 := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	merged := mergeTwoLists(l1, l2)
	for merged != nil {
		fmt.Println(merged.Val)
	}
}
