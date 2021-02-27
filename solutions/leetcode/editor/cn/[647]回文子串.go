package main
//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。 
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。 
//
// 
//
// 示例 1： 
//
// 输入："abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
// 
//
// 示例 2： 
//
// 输入："aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa" 
//
// 
//
// 提示： 
//
// 
// 输入的字符串长度不会超过 1000 。 
// 
// Related Topics 字符串 动态规划 
// 👍 491 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	/*
			dp[i][j] : s[i~j]是否为回文
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			dp[0][0] = 0

		 	0 0 0 0
		    0 1 0 0
			0 0 1 0
			0 0 0 1
	*/
	if len(s)==0 || len(s) ==1 {
		return len(s)
	}
	dp := make([][]bool,0)
	for i:=0;i<=len(s);i++ {
		dpi := make([]bool,len(s)+1)
		dp = append(dp,dpi)
	}
	ret := 0
	ret += len(s)
	for i :=1 ;i<=len(s);i++ {
		dp[i][i] = true
	}
	for i:=len(s);i>=1;i--{
		for j:=i+1;j<=len(s);j++ {
			if s[i-1]==s[j-1] {
				if j-i == 1 {
					dp[i][j] = true
				}else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j]  {
					ret ++
				}
			}
		}
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
