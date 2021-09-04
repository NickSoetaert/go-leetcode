package linked_lists

//https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
     Val int
     Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var runner = &ListNode{}
	runnerHead := runner

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			runner.Next = l1
			l1 = l1.Next
		} else {
			runner.Next = l2
			l2 = l2.Next
		}
		runner = runner.Next
	}

	if l1 == nil {
		runner.Next = l2
	} else {
		runner.Next = l1
	}
	return runnerHead.Next
}

