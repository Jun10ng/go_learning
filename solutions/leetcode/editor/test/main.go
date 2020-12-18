package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

// DP 解法
/*
	dp[i] = max(1,dp[i-1]*nums[i]) // 正数值最大
	neg_dp[i]                            // 数值最大
*/
func maxProduct(nums []int) (ans int) {
	imax := 1 //最大值
	imin := 1 //最小值 ，双负数情况
	max := math.MinInt64

	for _, e := range nums {
		if e < 0 {
			tmp := imax
			imax = imin
			imin = tmp
		}
		imax = int(math.Max(float64(e), float64(imax*e)))
		imin = int(math.Min(float64(e), float64(imin*e)))
		max = int(math.Max(float64(max), float64(imax)))
	}
	return max
}
