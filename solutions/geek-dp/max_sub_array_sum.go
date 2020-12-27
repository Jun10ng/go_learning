package main

import "fmt"

// 最大子序列问题

// 难点在于：dp[i]的定义是：以str[i]字符结尾的最大子数组和是多少

/*

示例：

输入：[-2, 1, -3, 4, -1, 3, -5, 1, 2]
输出：6
解释：连续子数组 [4,-1, 3] 的和最大为 6。
*/
func main() {
	fmt.Println(maxSubArraySum([]int{-2, 1, -3, 4, -1, 3, -5, 1, 2}))
}

func maxSubArraySum(nums []int) int {
	ans := -99990
	dp := make([]int, 0, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp = append(dp, -9999)
	}
	for i := 1; i < len(nums)+1; i++ {
		num := nums[i-1]
		dp[i] = max(dp[i-1]+num, num)
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
