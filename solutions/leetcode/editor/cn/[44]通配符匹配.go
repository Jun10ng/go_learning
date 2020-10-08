package main
//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。 
//
// '?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
// 
//
// 两个字符串完全匹配才算匹配成功。 
//
// 说明: 
//
// 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。 
// 
//
// 示例 1: 
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。 
//
// 示例 2: 
//
// 输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
// 
//
// 示例 3: 
//
// 输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
// 
//
// 示例 4: 
//
// 输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
// 
//
// 示例 5: 
//
// 输入:
//s = "acdcb"
//p = "a*c?b"
//输出: false 
// Related Topics 贪心算法 字符串 动态规划 回溯算法 
// 👍 532 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

// todo 使用回溯法超时了，以后使用动态规划解决
func isMatch(s string, p string) bool {
	p = prehandle(p)
	trace := func(s,p string)bool {
		return false
	}
	trace = func(s, p string) bool {
		if len(p)==0 {
			return len(s)==0
		}
		flag := len(s)!=0 && (s[0]==p[0] || p[0]=='?')
		if p[0]=='*' {
			if len(s)==0 {
				return isAllStart(p)
			}
			return trace(s[1:],p)|| trace(s,p[1:])
		}else {
			return flag && trace(s[1:],p[1:])
		}
	}
	return trace(s,p)
}
func isAllStart(s string)bool {
	for _,e := range s{
		if e!='*' {
			return false
		}
	}
	return true
}
// 去掉多余的*号
func prehandle(p string) string {
	ans := ""
	l := len(p)
	for i:=0;i<l;i++{
		if i!=l-1 &&p[i]=='*'&&p[i+1]=='*' {
			continue
		}
		ans = ans+string(p[i])
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
