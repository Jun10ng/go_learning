package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
func lengthOfLIS(nums []int) int {
	/*
			10,9,2,5,3,7,101,18
		   0 1 1 1 2 2 3  4  4
			dp[i] 代表为第 i 个数字为结尾的最长上升子序列的长度
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
