package main

import "fmt"

func main() {
	r := &TreeNode{Val: 8}
	r1 := &TreeNode{Val: 4}
	r2 := &TreeNode{Val: 3}
	r3 := &TreeNode{Val: 5}
	r4 := &TreeNode{Val: 2}
	r5 := &TreeNode{Val: 9}

	r.Left, r.Right = r1, r2
	r1.Left, r1.Right, r2.Left = r3, r4, r5
	//fmt.Println(generateParenthesis(3))

	fmt.Println(hammingWeight(00000000000000000000000000001011))

}

func hammingWeight(num uint32) int {
	ret := 0
	for num != 0 {
		num = num & (num - 1)
		ret++
	}
	return ret
}

// 上下左右
var vct = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}

func exist(board [][]byte, word string) bool {
	b := board
	ret := false
	trace := func(int, int, string) {}
	trace = func(i, j int, s string) {
		if len(s) == 0 {
			ret = true
			return
		}
		if i < 0 || j < 0 || i >= len(b) || j >= len(b[0]) {
			return
		}
		if b[i][j] != s[0] {
			return
		}
		var tmp byte
		tmp, b[i][j] = b[i][j], '#'
		for _, v := range vct {
			trace(i+v[0], j+v[1], s[1:])
			if ret {
				return
			}
		}
		b[i][j] = tmp
	}
	for i, ei := range b {
		for j, _ := range ei {
			trace(i, j, word)
			if ret {
				return ret
			}
		}
	}
	//trace(0,0,word)
	return ret
}

var penum = []string{"(", ")"}

func generateParenthesis(n int) []string {
	ret := []string{}

	/*
		每生成(, f+1,)f-- 如果f<0 说明目前的路径是无效的
		每生成一个括号，c++ 当c == 2*n 说明完毕
	*/

	trace := func(f, c int, cur string) {}
	trace = func(f, c int, cur string) {
		if f < 0 || f > n || c > 2*n {
			return
		}
		if c == 2*n && f == 0 {
			ret = append(ret, cur)
			return
		}
		for i, e := range penum {
			if i == 0 {
				trace(f+1, c+1, cur+e)
			} else {
				trace(f-1, c+1, cur+e)
			}
		}
	}
	trace(0, 0, "")
	return ret
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
