package main
//输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。 
//
// 
//
// 示例: 
//给定如下二叉树，以及目标和 sum = 22， 
//
//               5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
// 
//
// 返回: 
//
// [
//   [5,4,11,2],
//   [5,8,4,5]
//]
// 
//
// 
//
// 提示： 
//
// 
// 节点总数 <= 10000 
// 
//
// 注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/ 
// Related Topics 树 深度优先搜索 
// 👍 79 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	ans := &[][]int{}
	e := []int{}
	trace(root,0,sum,ans,e)
	return *ans
}

func trace(root *TreeNode,cur, sum int,ans *[][]int, e []int){
	// failure
	if root==nil {
		return
	}

	cur= cur + root.Val
	e = append(e,root.Val)
	// success
	if(sum == cur&&sum == cur && root.Left == nil && root.Right==nil){
		// 确保是叶节点
		cp := make([]int,len(e))
		copy(cp,e)
		*ans = append(*ans,cp)
		return
	}


	trace(root.Left,cur,sum,ans,e)
	trace(root.Right,cur,sum,ans,e)
}
//leetcode submit region end(Prohibit modification and deletion)
