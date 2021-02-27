package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubstrings("aaaaa"))
}
func countSubstrings(s string) int {
	/*
			dp[i][j] : s[i~j]是否为回文
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			dp[0][0] = 0

		 	0 0 0 0 0 0
		    0 1 0 0 0 0
			0 0 1 0 0 0
			0 0 0 1 1 1
			0 0 0 0 1 1
		0 0 0 0 0 1
	*/
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	dp := make([][]bool, 0)
	for i := 0; i <= len(s); i++ {
		dpi := make([]bool, len(s)+1)
		dp = append(dp, dpi)
	}
	ret := 0
	ret += len(s)
	for i := 1; i <= len(s); i++ {
		dp[i][i] = true
	}
	for i := len(s); i >= 1; i-- {
		for j := i + 1; j <= len(s); j++ {
			if s[i-1] == s[j-1] {
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] {
					ret++
				}
			}
		}
	}
	return ret
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
