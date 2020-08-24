package main
//输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。 
//
// 
//
// 参考以下这颗二叉搜索树： 
//
//      5
//    / \
//   2   6
//  / \
// 1   3 
//
// 示例 1： 
//
// 输入: [1,6,3,2,5]
//输出: false 
//
// 示例 2： 
//
// 输入: [1,3,2,6,5]
//输出: true 
//
// 
//
// 提示： 
//
// 
// 数组长度 <= 1000 
// 
// 👍 98 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
//后序遍历的根节点是最后一位
// 从第一个比根节点大的值到最后一位，都是右子数

func verifyPostorder(postorder []int) bool {
	return verify(postorder,0,len(postorder)-1)
}

func verify(p []int, s int, e int) bool{
	if s >= e {
		return true
	}
	root := p[e]
	i:= s
	for p[i]<root{
		i++
	}
	tmp :=i
	//从i开始的，有比root小的值，是不合理的。
	for tmp < e{
		if p[tmp] < root {
			return false
		}
		tmp++
	}
	return verify(p,s,i-1)&&verify(p,i,e-1)
}
//leetcode submit region end(Prohibit modification and deletion)
