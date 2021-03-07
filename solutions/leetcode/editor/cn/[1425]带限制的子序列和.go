package main
//给你一个整数数组 nums 和一个整数 k ，请你返回 非空 子序列元素和的最大值，子序列需要满足：子序列中每两个 相邻 的整数 nums[i] 和 num
//s[j] ，它们在原数组中的下标 i 和 j 满足 i < j 且 j - i <= k 。 
//
// 数组的子序列定义为：将数组中的若干个数字删除（可以删除 0 个数字），剩下的数字按照原本的顺序排布。 
//
// 
//
// 示例 1： 
//
// 输入：nums = [10,2,-10,5,20], k = 2
//输出：37
//解释：子序列为 [10, 2, 5, 20] 。
// 
//
// 示例 2： 
//
// 输入：nums = [-1,-2,-3], k = 1
//输出：-1
//解释：子序列必须是非空的，所以我们选择最大的数字。
// 
//
// 示例 3： 
//
// 输入：nums = [10,-2,-10,-5,20], k = 2
//输出：23
//解释：子序列为 [10, -2, -5, 20] 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= k <= nums.length <= 10^5 
// -10^4 <= nums[i] <= 10^4 
// 
// Related Topics 动态规划 
// 👍 60 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func constrainedSubsetSum(nums []int, k int) int {
	dp := make([]int,len(nums)+1)
	ret := nums[0]
	for i := 1;i<=len(nums);i++ {
		dp[i] = nums[i-1]
		for j:=i-1;i-j<=k&&j>0;j-- {
			dp[i] = max(dp[i],nums[i-1]+dp[j])
		}
		ret = max(ret,dp[i])
	}
	return ret
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
