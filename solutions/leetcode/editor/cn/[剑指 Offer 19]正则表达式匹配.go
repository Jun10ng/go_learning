package main

//请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配
//是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。 
//
// 示例 1: 
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
// 
//
// 示例 2: 
//
// 输入:
//s = "aa"
//p = "a*"
//输出: true
//解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
// 
//
// 示例 3: 
//
// 输入:
//s = "ab"
//p = ".*"
//输出: true
//解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
// 
//
// 示例 4: 
//
// 输入:
//s = "aab"
//p = "c*a*b"
//输出: true
//解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
// 
//
// 示例 5: 
//
// 输入:
//s = "mississippi"
//p = "mis*is*p*."
//输出: false 
//
// 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。 
// 
//
// 注意：本题与主站 10 题相同：https://leetcode-cn.com/problems/regular-expression-matching/
// 
// Related Topics 动态规划 
// 👍 85 👎 0
//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	if (p==""){
		if s == "" {
			return true
		}
		return false
	}
	var i int
	m := len(s)
	n := len(p)
	dp := make([][]bool, 0)
	for i:=0;i<m+1;i++ {
		dpi := make([]bool, n+1)
		dp = append(dp, dpi)
	}
	//dp[i][j]表示的是s的前i个字符和p的前j个字符匹配的结果。
	dp[0][0] = true
	for i=0;i<n;i++{
		if string(p[i])=="*" &&dp[0][i-1]{
			//dp[i][j]的i，j代表前i，j个字符，所以+1
			dp[0][i+1] = true
		}
	}
	for i=0;i<m;i++{
		for j:=0;j<n;j++ {
			if p[j]==s[i]||string(p[j])=="." {
				dp[i+1][j+1] = dp[i][j]
			}else if string(p[j])=="*" {
				//p的第j个字符和s的第i+1个字符不能匹配，
				//s="abc"，p="abcd*",让p的4,5个字符消失
				//只判断p的前j-1 和 s的 前i+1
				if string(p[j-1])==string(s[i]) || string(p[j-1])=="." {
					// 如果p的前一个字符是 . 或与s[i]相等（即p[j]字符出现0次的情况），取上一行的处理结果
					dp[i+1][j+1] = dp[i][j+1]
				}
				//如果是*，则取前两个字符的判断结果
				//p的第j个字符和s的第i+1个字符匹配结果
				dp[i+1][j+1] = dp[i+1][j+1]||dp[i+1][j-1]
			}
		}
	}
	return dp[m][n]
}
//leetcode submit region end(Prohibit modification and deletion)
