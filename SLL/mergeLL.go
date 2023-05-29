package sll

type Node struct {
	Val  int
	Next *Node
}

func MergeTwoLists(l1, l2 *Node) *Node {
	dummy := &Node{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}
	return dummy.Next
}
