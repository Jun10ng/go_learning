package main
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。 
//
// 
//
// 示例： 
//二叉树：[3,9,20,null,null,15,7], 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回其层序遍历结果： 
//
// 
//[
//  [3],
//  [9,20],
//  [15,7]
//]
// 
// Related Topics 树 广度优先搜索 
// 👍 812 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	ret := [][]int{}
	q := &queue{}
	q.push(root)
	for len(*q)!=0{
		reti := make([]int,0,len(*q))
		size := len(*q)
		for i:=0;i<size;i++{
			e := (*q)[0]
			reti = append(reti,e.Val)
			if e.Left!=nil{
				q.push(e.Left)
			}
			if e.Right!=nil {
				q.push(e.Right)
			}
			_ =q.pop()
		}
		ret = append(ret,reti)
	}
	return ret
}
type queue []*TreeNode

func (q *queue)push(t *TreeNode)  {
	*q = append(*q,t)
}
func (q *queue)pop() *TreeNode  {
	e := (*q)[0]
	(*q) = (*q)[1:]
	return e
}

//leetcode submit region end(Prohibit modification and deletion)
