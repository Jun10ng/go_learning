package main
//给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。 
//
// 
//
// 示例 1: 
//输入: 
//
// "bbbab"
// 
//
// 输出: 
//
// 4
// 
//
// 一个可能的最长回文子序列为 "bbbb"。 
//
// 示例 2: 
//输入: 
//
// "cbbd"
// 
//
// 输出: 
//
// 2
// 
//
// 一个可能的最长回文子序列为 "bb"。 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 只包含小写英文字母 
// 
// Related Topics 动态规划 
// 👍 354 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

//难点：注意这道题的计算方向
//
func longestPalindromeSubseq(s string) int {

	// bbabbbab
	// if s[i] = s[j] dp[i][j] = 2+dp[i+1][j-1]
	// if s[i]!= s[j] dp[i][j] = max{dp[i][j-1],dp[i+1][j]}

	dp := make([][]int,0)
	for i := 0; i < len(s)+1; i++ {
		dpi := make([]int,len(s)+1)
		dp = append(dp,dpi)
	}

	for i := 1; i < len(s)+1; i++ {
		dp[i][i]=1
	}

	for i := len(s)-1; i >= 0 ; i-- {
		for j := i+1; j < len(s); j++ {
			//for j := 1; j <len(s) ; j++ {
			//	for i := 0; i < j; i++ {
			switch  {
			case s[i]==s[j]:
				dp[i][j] = 2+dp[i+1][j-1]
			default:
				dp[i][j] = max(dp[i+1][j],dp[i][j-1])
			}

		}
	}
	return dp[0][len(s)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
