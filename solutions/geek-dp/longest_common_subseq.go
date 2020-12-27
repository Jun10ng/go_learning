package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] ：text1[0:i]和 text[0:j]的最长公共子序列长度
	// if text1[i] = text[j] dp[i][j] = 1+dp[i-1][j-1]
	// if text1[i]!= text[j] dp[i][j] = max{dp[i-1][j],dp[i][j-1]}

	//text1 = "abcde",
	//text2 = "ace"

	dp := make([][]int, 0, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dpi := make([]int, len(text2)+1)
		dp = append(dp, dpi)
	}

	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			c1, c2 := text1[i-1], text2[j-1]
			switch {
			case c1 == c2:
				dp[i][j] = 1 + dp[i-1][j-1]
			default:
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
