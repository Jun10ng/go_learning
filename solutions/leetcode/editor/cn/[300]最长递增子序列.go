package main
//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。 
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序
//列。 
// 
//
// 示例 1： 
//
// 
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1,0,3,2,3]
//输出：4
// 
//
// 示例 3： 
//
// 
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2500 
// -104 <= nums[i] <= 104 
// 
//
// 
//
// 进阶： 
//
// 
// 你可以设计时间复杂度为 O(n2) 的解决方案吗？ 
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗? 
// 
// Related Topics 二分查找 动态规划 
// 👍 1246 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	/*
		dp[i] 从起始位置到i元素的最长上升序列
	*/
	//numOfLIS,maxLen := 0,1
	maxLen := 1
	dp := make([]int,len(nums))
	// 表示以第 i 个数字结尾的序列的最长上升子序列的数量。
	for i,_:=range dp{
		dp[i] = 1
	}

	for j := 0; j < len(nums) ; j++ {
		for i := 0; i < j ; i++ {
			if nums[i] < nums[j] {
				dp[j] = max(dp[j],dp[i]+1)
			}
		}
		maxLen = max(dp[j],maxLen)
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
