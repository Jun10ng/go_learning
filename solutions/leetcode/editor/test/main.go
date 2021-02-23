package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{1}))
}
func maxSubArray(nums []int) int {
	/*
	  dp[i] 第i个元素结尾的最大子序和
	*/
	MinValue := -999999
	ret := MinValue
	dp := make([]int, len(nums)+1)
	for i, _ := range dp {
		dp[i] = MinValue
	}

	for i := 1; i < len(dp); i++ {
		num := nums[i-1]
		dp[i] = max(dp[i-1]+num, num)
		ret = max(dp[i], ret)
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	tmp := min2(a, b)
	return min2(tmp, c)
}
