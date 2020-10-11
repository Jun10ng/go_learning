package main
//给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换
//需遵循如下规则： 
//
// 
// 每次转换只能改变一个字母。 
// 转换后得到的单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回一个空列表。 
// 所有单词具有相同的长度。 
// 所有单词只由小写字母组成。 
// 字典中不存在重复的单词。 
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。 
// 
//
// 示例 1: 
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出:
//[
//  ["hit","hot","dot","dog","cog"],
//  ["hit","hot","lot","log","cog"]
//]
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。 
// Related Topics 广度优先搜索 数组 字符串 回溯算法 
// 👍 338 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 运行时间超时了
// 优化方案：两头回溯：由于知道了开始节点和结束节点，我们可以从两头都开始回溯，选择"路径"（差一个字母的单词）少的一边回溯。
// 直到两头节点有重合，说明连上了。
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// prehandle
	isEndWordExsit := false
	for _,e := range wordList{
		if e==endWord {
			isEndWordExsit = true
			break
		}
	}
	if !isEndWordExsit {
		return [][]string{}
	}

	//
	ans := [][]string{}
	minAnsLen := 1000000000
	rest := make(map[string]bool,len(wordList))
	for _,e := range wordList {
		rest[e] = true
	}
	trace := func(rest map[string]bool,cur []string,word string) {}
	trace = func(rest map[string]bool,cur []string,word string) {
		if len(cur)==0 {
			cur = append(cur,word)
		}
		if len(cur)>0 && cur[len(cur)-1]==endWord&&len(cur)<=minAnsLen {
			if  len(cur)< minAnsLen {
				minAnsLen = len(cur)
				//删除再添加
				ans =ans[:0]
			}
			cp := make([]string,len(cur))
			copy(cp,cur)
			ans = append(ans,cp)
			return

		}
		//
		legalWordList := getLegalWordList(word,rest)
		for _,legalWord := range legalWordList{
			// 在rest中删掉当前word
			rest[legalWord] = false
			cur = append(cur,legalWord)
			trace(rest,cur,legalWord)
			rest[legalWord] = true
			cur = cur[:len(cur)-1]
		}
	}
	trace(rest,[]string{},beginWord)
	return ans
}
func getLegalWordList(word string,rest map[string]bool) []string{
	ans := make([]string,0)
	for restWord,isExist := range rest{
		if !isExist {
			continue
		}
		if isLegal(word,restWord) {
			ans = append(ans,restWord)
		}
	}
	return ans
}
func isLegal(w,rw string)bool  {
	if len(w)!=len(rw){return false}
	onlyOneDiff := 0
	for i:=0;i<len(w)&&onlyOneDiff<=1;i++ {
		if w[i]!=rw[i] {
			onlyOneDiff++
		}
	}
	return onlyOneDiff==1
}
//leetcode submit region end(Prohibit modification and deletion)
