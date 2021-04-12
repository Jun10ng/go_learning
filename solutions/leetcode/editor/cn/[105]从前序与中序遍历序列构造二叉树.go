package main
//根据一棵树的前序遍历与中序遍历构造二叉树。 
//
// 注意: 
//你可以假设树中没有重复的元素。 
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
// Related Topics 树 深度优先搜索 数组 
// 👍 983 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	build := func(preorder []int, inorder []int) *TreeNode{
		return nil
	}
	build = func(pre []int, in []int) *TreeNode{
		if len(pre)==0 {
			return nil
		}
		r := &TreeNode{
			Val: pre[0],
		}
		if len(in)==1 {
			return r
		}
		idx := findIdx(in,pre[0]) // 左子树+根节点 节点数
		r.Left = build(pre[1:idx+1],in[:idx])
		r.Right = build(pre[idx+1:],in[idx+1:])
		return r
	}
	return build(preorder,inorder)
}
func findIdx(s []int,e int) int{
	for i,ie := range s{
		if ie == e{
			return i
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
