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
	//输入：nums = [10,9,2,5,3,7,101,18]
	//输出：4
	//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
	//{0,1,0,3,2,3} //  {0 1 2 3} 4

	/*
			dp[i] 以第i个元素结尾的最长递增子序列
		    决策 ： 第i个元素 放或不放
		    dp[i] = max(dp[i-1]+1, dp[j]+1) j 是上一个小于i元素的
	*/
	if len(nums)==1 {
		return 1
	}

	dp := make([]int,len(nums)+1)

	// begin
	findLastLess := func(i int) int{
		ans := -1
		for j:=i-1; j>-1;j-- {
			if nums[j]< nums[i] {
				ans = max(dp[j],ans)
			}
		}
		return ans
	}
	// end
	dp[0] = 1
	dp[1] = 1
	ans := 0
	for i:=1;i<len(nums);i++ {
		switch  {
		case nums[i]<nums[i-1]:
			dp[i] = max(findLastLess(i)+1,1)
		case nums[i]>nums[i-1]:
			//dp[i] = dp[i-1] + 1
			dp[i] = max(findLastLess(i)+1,dp[i-1]+1)
		default:
			dp[i] = dp[i-1]
		}
		ans = max(ans,dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
