package main
//反转一个单链表。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
// Related Topics 链表 
// 👍 1560 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	/*
		第一步，存下 下一个节点
		第二步 把当前节点指针指向上一个节点
		第三步 把上一个节点更新为当前节点，把当前节点更新为下一个节点，
	*/
	var pre *ListNode = nil
	var nxt *ListNode = nil
	cur := head
	for cur!=nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
//leetcode submit region end(Prohibit modification and deletion)
