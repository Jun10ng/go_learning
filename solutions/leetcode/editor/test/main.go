package main

import (
	"fmt"
)

func main() {
	fmt.Println(constrainedSubsetSum([]int{10, -2, -10, -5, 20}, 2))
	// 输入：nums = [10,-2,-10,-5,20], k = 2
	//输出：23
	//解释：子序列为 [10, -2, -5, 20] 。
}
func constrainedSubsetSum(nums []int, k int) int {
	dp := make([]int, len(nums)+1)
	ret := 0
	for i := 1; i <= len(nums); i++ {
		dp[i] = nums[i-1]
		for j := i - 1; i-j <= k && j > 0; j-- {
			dp[i] = max(dp[i], nums[i-1]+dp[j])
		}
		ret = max(ret, dp[i])
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
