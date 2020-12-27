package main
//给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。 
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。 
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
// 
//
// 若这两个字符串没有公共子序列，则返回 0。 
//
// 
//
// 示例 1: 
//
// 输入：text1 = "abcde", text2 = "ace" 
//输出：3  
//解释：最长公共子序列是 "ace"，它的长度为 3。
// 
//
// 示例 2: 
//
// 输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc"，它的长度为 3。
// 
//
// 示例 3: 
//
// 输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0。
// 
//
// 
//
// 提示: 
//
// 
// 1 <= text1.length <= 1000 
// 1 <= text2.length <= 1000 
// 输入的字符串只含有小写英文字符。 
// 
// Related Topics 动态规划 
// 👍 315 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] ：text1[0:i]和 text[0:j]的最长公共子序列长度
	// if text1[i] = text[j] dp[i][j] = 1+dp[i-1][j-1]
	// if text1[i]!= text[j] dp[i][j] = max{dp[i-1][j],dp[i][j-1]}

	//text1 = "abcde",
	//text2 = "ace"

	dp := make([][]int,0,len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dpi := make([]int,len(text2)+1)
		dp = append(dp,dpi)
	}

	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			c1,c2 :=text1[i-1],text2[j-1]
			switch {
			case c1 == c2:
				dp[i][j] = 1 + dp[i-1][j-1]
			default:
				dp[i][j] = max(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
