package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}))
	// 1 3 10 7 12
	// 1 2 6 11 9

}
func minSwap(A []int, B []int) int {
	/*
		第0状态 不交换
		第1状态 交换
		dp[i][j] 第i个元素，在第j状态下的最小交换次数
		dp[0][0] = 0 dp[0][1] = 1

		第i个元素换不换 取决于 i-1个元素是否有序
		状态转移方程
		dp[i][j] =
		A[i-1] < A[i] && B[i-1]<B[i]
			A[i-1] < B[i] && B[i-1] < A[i] //取最优值
				dp[i][0] = min(dp[i-1][0],dp[i-1][1]) //不交换
				dp[i][1] = min(dp[i-1][0],dp[i-1][1]) + 1 // 交换
			else // 这种情况一旦交换了 i元素 i-1元素也得交换，才能保持严格递增
				dp[i][0] = dp[i-1][0]
				dp[i][1] = dp[i-1][1] + 1
		else // 需要交换
			dp[i][0] = dp[i-1][1] //不交换，则上个位置必须交换
			dp[i][1] = dp[i-1][0]+1          //交换，则上位置不交换

		  0 0
		0 0 1
		0
		0
		0
	*/

}
func sumOfSlice(s []int, i, j int) int {
	sum := s[j]
	for _, e := range s[i:j] {
		sum += e
	}
	return sum
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
