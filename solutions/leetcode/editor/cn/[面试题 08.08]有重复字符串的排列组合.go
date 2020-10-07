package main

import (
	"sort"
	"strings"
)

//有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。
//
// 示例1: 
//
//  输入：S = "qqe"
// 输出：["eqq","qeq","qqe"]
// 
//
// 示例2: 
//
//  输入：S = "ab"
// 输出：["ab", "ba"]
// 
//
// 提示: 
//
// 
// 字符都是英文字母。 
// 字符串长度在[1, 9]之间。 
// 
// Related Topics 回溯算法 
// 👍 26 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func permutation(S string) []string {
	ans := make([]string,0)
	S = sortString(S)
	trace := func(cur, rest string,begin int) {}
	trace = func(cur, rest string, begin int) {
		if len(cur) == len(S) {
			cp := cur
			ans = append(ans,cp)
			return
		}
		if len(cur)>len(S) {
			return
		}

		//
		for i:= 0;i<len(rest);i++ {
			if i>0 && rest[i]==rest[i-1] {
				continue
			}
			cur = cur + string(rest[i])
			rest = rest[:i]+rest[i+1:]
			trace(cur,rest,i+1)
			rest = rest[:i] + string(cur[len(cur)-1])+rest[i:]
			cur = cur[:len(cur)-1]
		}
	}
	trace("",S,0)
	return ans
}
func sortString(s string) string {
	ss := strings.Split(s,"")
	sort.Strings(ss)
	return strings.Join(ss,"")
}
//leetcode submit region end(Prohibit modification and deletion)
