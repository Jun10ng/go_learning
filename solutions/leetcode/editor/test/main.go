package main

import (
	"fmt"
)

func main() {
	//m := [][]int{
	//	{0, 0}, {1, 1}, {0, 0},
	//}
	fmt.Println(cuttingRope(10))
}
func cuttingRope(n int) int {
	// 输入: 10
	//输出: 36
	//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
	/*
		dp[i] : 长度i 最大乘积

		决策： 每个单元都有两个状态：另起一段 or 放进上一段

		dp[i] = max{dp[i], (i-j)*dp[j)} 0 < j <i
	*/
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			dlt := i - j
			x := max(dlt*dp[j], dlt*j)
			dp[i] = max(x, dp[i])
		}
	}
	fmt.Println(dp)
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
