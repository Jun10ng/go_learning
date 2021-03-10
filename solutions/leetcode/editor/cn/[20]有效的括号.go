package main
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "()"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "()[]{}"
//输出：true
// 
//
// 示例 3： 
//
// 
//输入：s = "(]"
//输出：false
// 
//
// 示例 4： 
//
// 
//输入：s = "([)]"
//输出：false
// 
//
// 示例 5： 
//
// 
//输入：s = "{[]}"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 104 
// s 仅由括号 '()[]{}' 组成 
// 
// Related Topics 栈 字符串 
// 👍 2223 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func isValid(s string) bool {
	stk := stack{
		"",
	}
	for i := 0;i<len(s);i++ {
		e := string(s[i])
		if isEqual(e, stk.top()) {
			stk.pop()
		}else {
			stk.push(e)
		}
	}
	return len(stk.S) == 0
}
type stack struct {
	S string
}

func (s *stack)top() string {
	if len(s.S) == 0 {
		return ""
	}
	return s.S[len(s.S)-1:]
}
func (s *stack)push(e string)  {
	s.S = s.S + e
}
func (s *stack)pop() string {
	e := s.top()
	s.S = s.S[:len(s.S)-1]
	return e
}
func isEqual(s,e string) bool {
	switch s {
	case "]":
		return e == "["
	case "}":
		return e == "{"
	case ")":
		return e=="("
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
