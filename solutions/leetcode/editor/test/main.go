package main

import (
	"fmt"
)

func main() {

	//fmt.Print(lenLongestFibSubseq([]int{1,2,3,4,5,6,7,8}))
	fmt.Print(lenLongestFibSubseq([]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50})) // 4 14 18 32 50
}
func lenLongestFibSubseq(arr []int) int {
	/*
		状态
		dp[i][j] 以 arr[i] arr[j] 结尾的最长
		dp[i][j] = max(dp[i][j],dp[k][i]+1) arr[k]+arr[i] = arr[j] k<i
	*/
	dp := [][]int{}
	for i := 0; i <= len(arr); i++ {
		dpi := make([]int, len(arr)+1)
		for ei, _ := range dpi {
			dpi[ei] = 2
		}
		dp = append(dp, dpi)
	}
	ret := 0
	/*
		for i:=0;i<len(arr)-1;i++{
			for j:=i+1;j<len(arr);j++ {
				for k:=0;k<i;k++{
					if arr[i]+arr[k] == arr[j] {
						dp[i+1][j+1] = max2(dp[i+1][j+1],dp[k+1][i+1]+1)
						ret = max2(ret,dp[i+1][j+1])
					}
				}
			}
		}
		优化为下方代码
	*/
	pathMap := make(map[int]int, len(arr)) // pathMap为arr的逆
	for i, e := range arr {
		pathMap[e] = i + 1
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if k, ok := pathMap[arr[j]-arr[i]]; ok && k < i+1 {
				dp[i+1][j+1] = max2(dp[i+1][j+1], dp[k][i+1]+1)
				ret = max2(ret, dp[i+1][j+1])
			}
		}
	}

	return ret

}
func max2(a, b int) int {
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
