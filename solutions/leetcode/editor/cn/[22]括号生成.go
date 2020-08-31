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
func generateParenthesis(n int) []string {
	ans := &[]string{}
	line := ""
	trace(ans,&line,n,0)
	return *ans
}
func trace(ans *[]string,line *string, n,idx int){
	if idx == n+n {
		cp := ""
		cp = *line
		*ans = append(*ans,cp)
		return
	}
	*line = *line + "("
	if Ok(*line,n) {
		trace(ans,line,n,idx+1)
	}
	*line = (*line)[:idx]
	//
	*line = *line + ")"
	if Ok(*line,n) {
		trace(ans,line,n,idx+1)
	}
}
func Ok(s string,n int) bool {
	left,right := 0,0
	for _,e := range s{
		if e=='('{
			left++
		}else {
			right++
		}
	}
	if(left>n || right >n){
		return false
	}

	if len(s) == n+n{
		if isOk(s){return true}
		return false
	}
	return true
}

func isOk(s string) bool {
	stack := list.New()
	i := 0
	stack.PushBack(s[i])
	i++
	for i<len(s) {
		if s[i]=='(' {
			stack.PushBack(s[i])
		}else{
			if stack.Len()>0 {
				stack.Remove(stack.Back())
			}else {
				return false
			}
		}
		i++
	}
	return stack.Len()==0
}

//leetcode submit region end(Prohibit modification and deletion)
