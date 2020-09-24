package main

import "strings"

//给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
//
// 
//
// 示例：
//输入：S = "a1b2"
//输出：["a1b2", "a1B2", "A1b2", "A1B2"]
//
//输入：S = "3z4"
//输出：["3z4", "3Z4"]
//
//输入：S = "12345"
//输出：["12345"]
// 
//
//
// 提示： 
//
// 
// S 的长度不超过12。 
// S 仅由数字和字母组成。 
// 
// Related Topics 位运算 回溯算法 
// 👍 209 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func letterCasePermutation(S string) []string {
	ans := make([]string,0)
	cur := make([]string,0)
	ss := make([]string,0)
	for _,e := range S{
		ss = append(ss,strings.ToLower(string(e)))
	}

	trace := func(cur []string,idx int) {}
	trace = func(cur []string, idx int) {
		if idx == len(ss) {
			cp := strings.Join(cur,"")
			ans = append(ans,cp)
			return
		}
		if ss[idx] <= "9" {
			cur = append(cur,ss[idx])
			trace(cur,idx+1)
		}else {
			cur = append(cur,ss[idx])
			trace(cur,idx+1)
			cur = cur[:len(cur)-1]

			cur = append(cur,strings.ToUpper(ss[idx]))
			trace(cur,idx+1)
			cur = cur[:len(cur)-1]
		}
	}
	trace(cur,0)

	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
