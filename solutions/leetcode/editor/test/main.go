package main

import (
	"fmt"
)

func main() {
	m := [][]int{
		{0, 0}, {1, 1}, {0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(m))
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < len(obstacleGrid); i++ {
		dpi := make([]int, len(obstacleGrid[0]))
		dp = append(dp, dpi)
	}
	// 初始化
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		}
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		}
	}
	// 决策
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			c := obstacleGrid[i][j]
			switch c {
			case 1:
				dp[i][j] = 0
			case 0:
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
