package main

import "container/list"

//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
//
// 
//
// 例如: 
//给定二叉树: [3,9,20,null,null,15,7], 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回： 
//
// [3,9,20,15,7]
// 
//
// 
//
// 提示： 
//
// 
// 节点总数 <= 1000 
// 
// Related Topics 树 广度优先搜索 
// 👍 33 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nq := list.New()
	ret := make([]int,0)

	nq.PushBack(root)
	for nq.Len()!=0 {
		n := nq.Front().Value.(*TreeNode)
		nq.Remove(nq.Front())

		ret = append(ret,n.Val)
		if n.Left!=nil {
			nq.PushBack(n.Left)
		}
		if n.Right!=nil {
			nq.PushBack(n.Right)
		}
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
