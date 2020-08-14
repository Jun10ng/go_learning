package main

import "fmt"

//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 5000
//
//
//
// 注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-
//preorder-and-inorder-traversal/
// Related Topics 树 递归
// 👍 174 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	for {

	}
}

func build(pre []int, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = pre[0]
	idxRoot := indexOf(in, root.Val)
	root.Left = build(pre[1:idxRoot+1], in[:idxRoot])
	root.Right = build(pre[idxRoot+1:], in[idxRoot+1:])
	return root
}

func indexOf(s []int, i int) int {
	for x, y := range s {
		if i == y {
			return x
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	t := buildTree(pre, in)
	fmt.Print(t)
}
