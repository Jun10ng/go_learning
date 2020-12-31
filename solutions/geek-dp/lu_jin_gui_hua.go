package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 3))
}
func uniquePaths(m int, n int) int {
	// dp[i][j] 从起点 到达i，j 有几种走法
	dp := make([][]int, 0)
	for i := 0; i < m+1; i++ {
		dpi := make([]int, n+1)
		dp = append(dp, dpi)
	}
	dp[0][1] = 1
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] += dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
