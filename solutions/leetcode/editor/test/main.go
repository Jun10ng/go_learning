package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDistinct("babgbag", "bag"))
}
func numDistinct(s string, t string) int {
	/*
			r a b b b i t
		  r 1 1 1 1 1 1 1
		  a	0 1 1 1 1 1 1
		  b 0 1 1 2 3 3 3
		  b     0 1 3 3 3
		  i         0 3 3
		  t
		dp[i][i] = dp[i-1][i-1]
		dp[i][j]
		s[i] == t[j]
			else dp[i-1][j-1] + dp[i][j-1]
		else dp[i][j] = dp[i][j-1]

			b a b g b a g
		  b 1 1 2 2 3 3 3
		  a   1 1 1 1 4 4
		  g     0 1 1 1 5
	*/
	dp := make([][]int, 0)
	for i := 0; i <= len(t); i++ {
		dpi := make([]int, len(s)+1)
		dp = append(dp, dpi)
	}
	for j := 0; j <= len(s); j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= len(t); i++ {
		for j := i; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
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
