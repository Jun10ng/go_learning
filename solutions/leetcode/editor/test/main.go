package main

import (
	"fmt"
)

func main() {
	r := &TreeNode{Val: 8}
	r1 := &TreeNode{Val: 4}
	r2 := &TreeNode{Val: 3}
	r3 := &TreeNode{Val: 5}
	r4 := &TreeNode{Val: 2}
	r5 := &TreeNode{Val: 9}
	r.Left = r1
	r.Right = r5
	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	fmt.Println(lowestCommonAncestor(r, r4, r3))
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	// 左右子树各找到一个 就返回root
	return root
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
