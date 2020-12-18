package main

import "math"

//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 
//
// 示例 1: 
//
// 输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
// 
//
// 示例 2: 
//
// 输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。 
// Related Topics 数组 动态规划 
// 👍 862 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// DP 解法
// DP 解法
/*
	dp[i] = max(1,dp[i-1]*nums[i]) // 正数值最大
	neg_dp[i]                            // 数值最大
*/
func maxProduct(nums []int) (ans int) {
	imax := 1; //最大值
	imin := 1; //最小值 ，双负数情况
	max := math.MinInt64;

	for _,e := range nums{
		if e<0 {
			tmp := imax
			imax = imin
			imin = tmp
		}
		imax = int(math.Max(float64(e),float64(imax*e)))
		imin = int(math.Min(float64(e),float64(imin*e)))
		max = int(math.Max(float64(max),float64(imax)))
	}
	return max
}


//leetcode submit region end(Prohibit modification and deletion)
