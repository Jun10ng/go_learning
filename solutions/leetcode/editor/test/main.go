package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

func isMatch(s string, p string) bool {
	if s == "" || p == "" {
		return false
	}
	var i int
	m := len(s)
	n := len(p)
	dp := make([][]bool, 0)
	for i := 0; i < m+1; i++ {
		dpi := make([]bool, n+1)
		dp = append(dp, dpi)
	}
	//dp[i][j]表示的是s的前i个字符和p的前j个字符匹配的结果。
	dp[0][0] = true
	for i = 0; i < n; i++ {
		if string(p[i]) == "*" && dp[0][i-1] {
			//dp[i][j]的i，j代表前i，j个字符，所以+1
			dp[0][i+1] = true
		}
	}
	for i = 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == s[i] || string(p[j]) == "." {
				dp[i+1][j+1] = dp[i][j]
			} else if string(p[j]) == "*" {
				//p的第j个字符和s的第i+1个字符不能匹配，
				//s="abc"，p="abcd*",让p的4,5个字符消失
				//只判断p的前j-1 和 s的 前i+1
				if string(p[j-1]) == string(s[i]) || string(p[j-1]) == "." {
					// 如果p的前一个字符是 . 或与s[i]相等（即p[j]字符出现0次的情况），取上一行的处理结果
					dp[i+1][j+1] = dp[i][j+1]
				}
				//如果是*，则取前两个字符的判断结果
				//p的第j个字符和s的第i+1个字符匹配结果
				dp[i+1][j+1] = dp[i+1][j+1] || dp[i+1][j-1]
			}
		}
	}
	return dp[m][n]
}
func main() {
	s := "ab"
	p := ".*"

	fmt.Print(isMatch(s, p))
}
