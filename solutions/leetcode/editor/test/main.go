package main

import (
	"fmt"
)

func main() {
	fmt.Println(completeBagV1([]int{0, 3, 2, 1}, []int{0, 5, 2, 3}, 3, 5))
	fmt.Println(completeBagV2([]int{0, 3, 2, 1}, []int{0, 5, 2, 3}, 3, 5))
}

func completeBagV1(w, n []int, N, W int) int {
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dpi := make([]int, W+1)
		dp[i] = dpi
	}

	// 遍历每一件商品
	for i := 1; i <= N; i++ {
		// 遍历背包容量
		for rw := 1; rw <= W; rw++ {
			//尽可能放入多次物品
			dp[i][rw] = dp[i-1][rw]
			for k := 0; k <= rw/w[i]; k++ {
				dp[i][rw] = max(dp[i][rw], dp[i-1][rw-k*w[i]]+k*n[i])
			}
		}
	}
	fmt.Println("============V1============")
	for _, dpi := range dp {
		fmt.Println(dpi)
	}
	return dp[N][W]
}
func completeBagV2(w, n []int, N, W int) int {
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dpi := make([]int, W+1)
		dp[i] = dpi
	}

	// 遍历每一件商品
	for i := 1; i <= N; i++ {
		// 遍历背包容量
		for rw := 1; rw <= W; rw++ {
			//尽可能放入多次物品
			if rw >= w[i] {
				dp[i][rw] = max(dp[i-1][rw], dp[i][rw-w[i]]+n[i])
			}
		}
	}
	fmt.Println("============V2============")
	for _, dpi := range dp {
		fmt.Println(dpi)
	}
	return dp[N][W]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
