package main
//给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。 
//
// 示例: 
//
// 
//输入: [4, 6, 7, 7]
//输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7
//]] 4
//
// 说明: 
//
// 
// 给定数组的长度不会超过15。 
// 数组中的整数范围是 [-100,100]。 
// 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。 
// 
// Related Topics 深度优先搜索 
// 👍 234 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findSubsequences(nums []int) [][]int {
	ans := [][]int{}
	doFind := func(path []int,i int){}
	doFind = func(path []int,i int) {
		if len(path)>1 {
			cpy := make([]int,len(path))
			copy(cpy,path)
			ans = append(ans,cpy)
		}
		vis := make(map[int]bool) // 同一层的递归入口不能重复，在第三层递归选了第一个7，就不能在第三次的递归再选第二个7
		for j := i;j<len(nums);j++ {
			if vis[nums[j]]{
				continue
			}
		if len(path)==0 || nums[j]>=path[len(path)-1] {
				vis[nums[j]] = true
				path = append(path,nums[j])
				doFind(path,j+1)
				path = path[:len(path)-1]
			}
		}
	}

	doFind([]int{},0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
