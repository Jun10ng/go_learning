package main
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。 
//
// 假设一个二叉搜索树具有如下特征： 
//
// 
// 节点的左子树只包含小于当前节点的数。 
// 节点的右子树只包含大于当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
// 
//
// 示例 1: 
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
// 
//
// 示例 2: 
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
// 
// Related Topics 树 深度优先搜索 递归 
// 👍 962 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//// 前序解法
//var pre int = 9999999999
//func isValidBST(root *TreeNode) bool {
//	// 前序遍历
//	pre = 9999999999
//	return helper(root,0-pre,pre)
//}
//func helper(root *TreeNode,lower,upper int) bool {
//	if root == nil {
//		return true
//	}
//	if root.Val<=lower || root.Val >= upper {
//		return false
//	}
//	return helper(root.Left,lower,root.Val) && helper(root.Right,root.Val,upper)
//}


// ==========
//// 中序解法 二叉排序树的中序遍历是递增数列
// var pre int = -9999999999
//func isValidBST(root *TreeNode) bool {
//	// 中序遍历 左节点->根节点->右节点
//	pre = -9999999999
//	return helper(root)
//}
//func helper(root *TreeNode) bool {
//	// 中序遍历 左节点->根节点->右节点
//	if root == nil {
//		return true
//	}
//	if !helper(root.Left) {
//		return false
//	}
//	if root.Val <= pre {
//		return false
//	}
//	pre = root.Val
//	return helper(root.Right)
//}
//leetcode submit region end(Prohibit modification and deletion)
