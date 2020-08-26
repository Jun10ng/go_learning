package main

import "container/list"

//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
// 返回其层次遍历结果： 
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
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
// 👍 40 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//* Definition for a binary tree node.
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	q := list.New()

	return doLevelOrder(root,ans,1,q)
}

func doLevelOrder(root *TreeNode,ans [][]int,level int,q *list.List) [][]int {
	if root == nil {
		return ans
	}
	q.PushBack(root)

	for q.Len()!=0 {
		v :=make([]int,0)
		len := q.Len()
		levelNodes := make([]TreeNode,0,len)
		for i:=0;i<len;i++ {
			e,_:= q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			levelNodes = append(levelNodes,*e)
			v = append(v,e.Val)

			if e.Right!=nil {
				q.PushBack(e.Right)
			}
			if e.Left != nil {
				q.PushBack(e.Left)
			}
		}
		if level%2 ==1 {
			v = reverse(v)
		}
		level++
		ans = append(ans,v)
	}
	return ans
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
//leetcode submit region end(Prohibit modification and deletion)
