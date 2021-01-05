package main

import (
	"fmt"
)

func main() {
	//m := []int{4, 6, 7, 7} // 1357 1347
	fmt.Println(climbStairs(4))
}
func climbStairs(n int) int {
	/*
			dp[i] 到达i位置有几种跳法
		    决策 使用1 跳 还是2 跳
		    1 | 1,2,1
	*/

	dp := make([]int, n+1)
	dp[0] = 1
	stepList := []int{1, 2}
	for i := 1; i < n+1; i++ {
		for _, step := range stepList {
			if i-step >= 0 {
				dp[i] += dp[i-step]
			}
		}
	}
	return dp[n]
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
