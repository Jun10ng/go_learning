package main
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。 
//
// 示例 1： 
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
// 
//
// 示例 2： 
//
// 输入: "cbbd"
//输出: "bb"
// 
// Related Topics 字符串 动态规划 
// 👍 3037 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s)<=1 {
		return s
	}
	dp := make([][]bool,0)
	ans := string(s[0])
	for i:=0;i<len(s)+1;i++ {
		dpi := make([]bool,len(s)+1)
		dp = append(dp,dpi)
	}

	for i:=1;i<len(dp);i++ {
		dp[i][i]=true
	}

	for j:=1;j<len(s);j++ {
		for i:=0;i<j;i++ {
			dp[i][j] = s[i]==s[j] && (j-i<3 || dp[i+1][j-1])
			if dp[i][j] && j-i + 1> len(ans) {
				ans = s[i:j+1]
			}
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
