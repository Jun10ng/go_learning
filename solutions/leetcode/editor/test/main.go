package main

func main() {
}

func wordBreak(s string, wordDict []string) bool {

	/*// 输入: s = "leetcode", wordDict = ["leet", "code"]
	  //输出: true
	  //解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。*/

	dp := make([]bool, len(s)+1)

	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}
	wordMap[""] = true
	dp[0] = true
	for j := 1; j < len(dp); j++ {
		for i := 0; i < j; i++ {
			str2 := s[i:j]
			if dp[i] && wordMap[str2] {
				dp[j] = true
				break
			}
		}
	}
	return dp[len(s)]
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
