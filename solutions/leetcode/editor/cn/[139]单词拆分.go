package main
//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。 
//
// 说明： 
//
// 
// 拆分时可以重复使用字典中的单词。 
// 你可以假设字典中没有重复的单词。 
// 
//
// 示例 1： 
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 
//
// 示例 2： 
//
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
// 
//
// 示例 3： 
//
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
// 
// Related Topics 动态规划 
// 👍 795 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {

	/*// 输入: s = "leetcode", wordDict = ["leet", "code"]
	  //输出: true
	  //解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。*/

	dp := make([]bool,len(s)+1)

	wordMap := make(map[string]bool,len(wordDict))
	for _,word:=range wordDict{
		wordMap[word] = true
	}
	wordMap[""]=true
	dp[0] = true
	for j:=1;j<len(dp);j++ {
		for i := 0; i <j; i++ {
			str2 := s[i:j]
			if dp[i] && wordMap[str2] {
				dp[j] = true
				break
			}
		}
	}
	return dp[len(s)]
}
//leetcode submit region end(Prohibit modification and deletion)
