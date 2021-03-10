package main

import "fmt"

func main() {
	root := &ListNode{
		Val: 0,
	}
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 3,
	}
	//n4 := &ListNode{
	//	Val: 4,
	//	Next: nil,
	//}
	root.Next = n1
	n1.Next, n2.Next, n3.Next = n2, n3, n1
	rvs := hasCycle(root)
	fmt.Print(rvs)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
