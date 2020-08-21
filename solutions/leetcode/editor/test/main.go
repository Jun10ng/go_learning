package main

import (
	"container/list"
)

//leetcode submit region begin(Prohibit modification and deletion)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	nq := list.New()
	ret := make([]int, 0)

	nq.PushBack(root)
	for nq.Len() != 0 {
		n := nq.Front().Value.(*TreeNode)
		nq.Remove(nq.Front())

		ret = append(ret, n.Val)
		if n.Left != nil {
			nq.PushBack(n.Left)
		}
		if n.Right != nil {
			nq.PushBack(n.Right)
		}
	}
	return ret
}
func main() {

}
