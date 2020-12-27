package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, 0)
	ans := string(s[0])
	for i := 0; i < len(s)+1; i++ {
		dpi := make([]bool, len(s)+1)
		dp = append(dp, dpi)
	}

	for i := 1; i < len(dp); i++ {
		dp[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
