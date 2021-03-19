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

	r.Left, r.Right = r1, r2
	r1.Left, r1.Right, r2.Left = r3, r4, r5
	fmt.Println(minDepth(r))
}
func minDepth(root *TreeNode) int {
	depth := 1
	if root.Right == nil && root.Left == nil {
		return depth
	}
	l, r := 99999999, 99999999
	if root.Left != nil {
		l = minDepth(root.Left)
	}
	if root.Right != nil {
		r = minDepth(root.Right)
	}

	if l < r {
		return l + 1
	}
	return r + 1
}
func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := [][]int{}
	q := &queue{}
	q.push(root)
	for len(*q) != 0 {
		reti := make([]int, 0, len(*q))
		size := len(*q)
		for i := 0; i < size; i++ {
			e := (*q)[0]
			reti = append(reti, e.Val)
			if e.Left != nil {
				q.push(e.Left)
			}
			if e.Right != nil {
				q.push(e.Right)
			}
			_ = q.pop()
		}
		ret = append(ret, reti)
	}
	return ret
}

type queue []*TreeNode

func (q *queue) push(t *TreeNode) {
	*q = append(*q, t)
}
func (q *queue) pop() *TreeNode {
	e := (*q)[0]
	(*q) = (*q)[1:]
	return e
}

func majorityElement(nums []int) int {
	/*
		{2,2,1,1,1,2,2}
		[2 2 1] [1 1 2 2]

	*/
	if len(nums) == 1 {
		return nums[0]
	}
	m := len(nums) / 2
	l := majorityElement(nums[:m])
	r := majorityElement(nums[m:])
	if l == r {
		return l
	}
	lc, rc := 0, 0
	for _, e := range nums {
		if e == l {
			lc++
		}
		if e == r {
			rc++
		}
	}
	if lc > rc {
		return l
	}
	return r
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x, n-1)
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
