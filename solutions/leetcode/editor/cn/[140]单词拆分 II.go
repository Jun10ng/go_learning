package main
//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的
//句子。 
//
// 说明： 
//
// 
// 分隔时可以重复使用字典中的单词。 
// 你可以假设字典中没有重复的单词。 
// 
//
// 示例 1： 
//
// 输入:
//s = "catsanddog"
//wordDict = ["cat", "cats", "and", "sand", "dog"]
//输出:
//[
//  "cats and dog",
//  "cat sand dog"
//]
// 
//
// 示例 2： 
//
// 输入:
//s = "pineapplepenapple"
//wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
//输出:
//[
//  "pine apple pen apple",
//  "pineapple pen apple",
//  "pine applepen apple"
//]
//解释: 注意你可以重复使用字典中的单词。
// 
//
// 示例 3： 
//
// 输入:
//s = "catsandog"
//wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出:
//[]
// 
// Related Topics 动态规划 回溯算法 
// 👍 247 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 为啥又是运行时间超时啊！！！
func wordBreak(s string, wordDict []string) []string {
	minWordLen := getMinWordLen(wordDict)
	ans := make([]string,0)
	wordMap := make(map[string]bool)
	for _,e := range wordDict{
		wordMap[e] = true
	}
	//
	cur := ""
	//isWordDone := false //判断是否刚找到一个单词，用与idx+1 还是 idx+minWordLen
	trace := func(start int){}
	trace = func(start int) {
		if start >= len(s) {
			cp := cur
			ans = append(ans,cp)
			return
		}
		//
		for i:=start+minWordLen-1;i<=len(s);i++{
			tmpWord := s[start:i]
			if _,ok := wordMap[tmpWord];ok {
				oldCur := cur
				if len(cur)==0 {
					cur = cur + tmpWord
				}else {
					cur = cur + " " + tmpWord
				}
				trace(i)
				cur = oldCur
			}
		}
	}
	trace(0)
	return ans
}
func getMinWordLen(wl []string)int  {
	minLen := 10000000
	for _,e := range wl{
		if minLen > len(e) {
			minLen = len(e)
		}
	}
	return minLen
}
//leetcode submit region end(Prohibit modification and deletion)
