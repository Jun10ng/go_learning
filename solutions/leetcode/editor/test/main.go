package main

import "fmt"

func main() {

	fmt.Println(numTilings(4))
}

func numTilings(N int) int {
	/*
		输入 3 输出  5

		dp[i] 为 搭建方案种数
		dp[0] = 0 dp[1] =1 dp[2]=2 dp[3] = 5

		dp[i] = max(dp[i-j] * dp[j]) 0<j<i
	*/

	dp := make([]int, N+1)
	dp[0], dp[1], dp[2], dp[3] = 0, 1, 2, 5

	for i := 4; i < N+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] += dp[i-j] * dp[j]
		}
	}

	return dp[N]
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
