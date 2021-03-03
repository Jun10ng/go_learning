package main
//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。 
//
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列
//，而 "AEC" 不是） 
//
// 题目数据保证答案符合 32 位带符号整数范围。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//(上箭头符号 ^ 表示选取的字母)
//rabbbit
//^^^^ ^^
//rabbbit
//^^ ^^^^
//rabbbit
//^^^ ^^^
// 
//
// 示例 2： 
//
//
//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
//(上箭头符号 ^ 表示选取的字母)
//babgbag
//^^ ^
//babgbag
//^^    ^
//babgbag
//^    ^^
//babgbag
//  ^  ^^
//babgbag
//    ^^^ 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length, t.length <= 1000 
// s 和 t 由英文字母组成 
// 
// Related Topics 字符串 动态规划 
// 👍 312 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
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
	dp := make([][]int,0)
	for i:=0;i<=len(t);i++{
		dpi := make([]int,len(s)+1)
		dp = append(dp,dpi)
	}
	for j:=0;j<=len(s);j++ {
		dp[0][j] = 1
	}

	for i:=1;i<=len(t);i++ {
		for j:=i;j<=len(s);j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			}else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
//leetcode submit region end(Prohibit modification and deletion)
