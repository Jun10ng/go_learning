package main

import "fmt"

// 回文串个数

/*
	给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
	具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

示例2：

输入："aaa"
输出：6
解释：共有六个回文子串，分别为 "a", "a", "a", "aa", "aa", "aaa"。
注意题设，具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串，因此像 "aa" 和 "aa" 就是两个不同的回文子串。
*/
func main() {
	fmt.Println(dp("aaa"))
}
func dp(str string) int {
	ans := 0
	dp := make([][]bool, 0)
	for i := 0; i < len(str)+1; i++ {
		dpi := make([]bool, len(str)+1)
		dp = append(dp, dpi)
	}

	for i := 1; i < len(str)+1; i++ {
		dp[i][i] = true
		ans++
	}

	for j := 1; j < len(str); j++ {
		for i := 0; i < j; i++ {
			// abccba
			dp[i][j] = str[i] == str[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] {
				ans++
			}
		}
	}

	return ans

}
