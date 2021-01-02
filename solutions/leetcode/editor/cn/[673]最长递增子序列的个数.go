package main
//给定一个未排序的整数数组，找到最长递增子序列的个数。 
//
// 示例 1: 
//
// 
//输入: [1,3,5,4,7]
//输出: 2
//解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
// 
//
// 示例 2: 
//
// 
//输入: [2,2,2,2,2]
//输出: 5
//解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
// 
//
// 注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。 
// Related Topics 动态规划 
// 👍 252 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findNumberOfLIS(nums []int) int {
	/*
		dp[i] 从起始位置到i元素的最长上升序列
		count[i] 从起始位置到i元素的最长上升序列个数
	*/
	numOfLIS,maxLen := 0,1
	dp := make([]int,len(nums))
	// 表示以第 i 个数字结尾的序列的最长上升子序列的数量。
	// 解题思路：相当于对count 做一次DP
	count := make([]int,len(nums))
	for i,_:=range dp{
		dp[i] = 1
		count[i] = 1
	}

	for j := 0; j < len(nums) ; j++ {
		for i := 0; i < j ; i++ {
			if nums[i] < nums[j] {
				switch {
				case dp[i]+1 > dp[j]:
					// 此时需要更新dp
					dp[j] = dp[i]+1
					count[j] = count[i] // 最长子序列在延续
				case dp[i]+1 == dp[j]:
					count[j] += count[i] //出现了长度相同的最长子序列，叠加
				}
			}
		}
		maxLen = max(dp[j],maxLen)
	}
	for i,e := range dp{
		if e==maxLen {
			numOfLIS+=count[i]
		}
	}
	return numOfLIS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
