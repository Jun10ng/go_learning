package main

import (
	"sort"
	"strings"
)

//输入一个字符串，打印出该字符串中字符的所有排列。
//
// 
//
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。 
//
// 
//
// 示例: 
//
// 输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
// 
//
// 
//
// 限制： 
//
// 1 <= s 的长度 <= 8 
// Related Topics 回溯算法 
// 👍 98 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func permutation(s string) []string {
	ans := &[]string{}
	var word string

	// 解决重复字符
	s = sortString(s)
	n := len(s)
	used := make([]bool,n)
	i := 0
	trace(ans,word,used,s,n,i)
	return *ans
}

func trace(ans *[]string, word string,used []bool,s string ,n int,idx int)  {
	if idx==n {
		var cp string
		cp = word
		*ans = append(*ans,cp)
		return
	}

	for i:=0;i<n;i++{
		if i>0 {
			if used[i-1]&&s[i]==s[i-1] {
				continue
			}
		}
		if !used[i] {
			used[i] = true
			word = word+string(s[i])
			trace(ans,word,used,s,n,idx+1)
			word = word[:len(word)-1]
			used[i] = false
		}
	}

}

func sortString(s string)string{
	ts := strings.Split(s,"")
	sort.Strings(ts)
	return strings.Join(ts,"")
}

//leetcode submit region end(Prohibit modification and deletion)
