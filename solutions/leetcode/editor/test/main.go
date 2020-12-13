package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(coinChange([]int{2, 3, 5}, 12))
}

var MAX = 99999

// DP 解法
func coinChange(coins []int, amount int) int {
	dp := make([]int, 0)
	for i := 0; i <= amount; i++ {
		dp = append(dp, MAX)
	}
	// 初始化状态
	dp[0] = 0

	// 状态
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i < coin {
				continue
			}
			// 决策
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
		}
	}
	return dp[amount]
}
