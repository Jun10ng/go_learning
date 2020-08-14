package main


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
	Val int
	Left *TreeNode
	Right *TreeNode
}


//leetcode submit region begin(Prohibit modification and deletion)
//法二：非递归使用栈


func indexOf(s []int,i int) int {
	for x,y := range s{
		if i==y{
			return x
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)

//法一：递归
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder)==0||len(inorder)==0 {
//		return nil
//	}
//	root := new(TreeNode)
//	root.Val = preorder[0]
//	idxRoot :=indexOf(inorder,root.Val)
//	root.Left = buildTree(preorder[1:idxRoot+1],inorder[:idxRoot])
//	root.Right = buildTree(preorder[idxRoot+1:],inorder[idxRoot+1:])
//	return root
//}