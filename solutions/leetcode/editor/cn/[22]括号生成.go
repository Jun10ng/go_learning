package main

import "container/list"

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 
//
// 示例： 
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
// 
// Related Topics 字符串 回溯算法 
// 👍 1274 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var penum = []string{"(",")"}
func generateParenthesis(n int) []string {
	ret := []string{}

	/*
		每生成(, f+1,)f-- 如果f<0 说明目前的路径是无效的
		每生成一个括号，c++ 当c == 2*n 说明完毕
	*/

	trace := func(f,c int,cur string) {}
	trace = func(f,c int,cur string) {
		if f<0 || f >n || c>2*n{
			return
		}
		if c == 2*n && f==0 {
			ret  = append(ret,cur)
			return
		}
		for i,e := range penum {
			if i==0{
				trace(f+1,c+1,cur+e)
			}else {
				trace(f-1,c+1,cur+e)
			}
		}
	}
	trace(0,0,"")
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
