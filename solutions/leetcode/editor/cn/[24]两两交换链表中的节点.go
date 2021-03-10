package main
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。 
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 
//
// 
//
// 示例 1： 
//
// 
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
// 
//
// 示例 2： 
//
// 
//输入：head = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：head = [1]
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点的数目在范围 [0, 100] 内 
// 0 <= Node.val <= 100 
// 
//
// 
//
// 进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。） 
// Related Topics 递归 链表 
// 👍 847 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	/*
		1. 存下pre，遍历两个节点，存下nxt
		2、 pre->n2 n2->n1 n1->nxt
		3  pre = n1
	*/
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	var n2 *ListNode = nil
	var nxt *ListNode = nil
	var newHeade *ListNode = nil
	n1 := head
	for n1!=nil {
		n2 = n1.Next
		if n2 == nil {
			break
		}
		nxt = n2.Next
		if pre!=nil {
			pre.Next = n2
		}else {
			newHeade = n2
		}
		n2.Next = n1
		n1.Next = nxt
		pre = n1
		n1 = nxt
	}
	return newHeade
}
//leetcode submit region end(Prohibit modification and deletion)
