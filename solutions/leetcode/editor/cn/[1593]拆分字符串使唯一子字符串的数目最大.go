package main
//给你一个字符串 s ，请你拆分该字符串，并返回拆分后唯一子字符串的最大数目。 
//
// 字符串 s 拆分后可以得到若干 非空子字符串 ，这些子字符串连接后应当能够还原为原字符串。但是拆分出来的每个子字符串都必须是 唯一的 。 
//
// 注意：子字符串 是字符串中的一个连续字符序列。 
//
// 
//
// 示例 1： 
//
// 输入：s = "ababccc"
//输出：5
//解释：一种最大拆分方法为 ['a', 'b', 'ab', 'c', 'cc'] 。像 ['a', 'b', 'a', 'b', 'c', 'cc'] 这样
//拆分不满足题目要求，因为其中的 'a' 和 'b' 都出现了不止一次。
// 
//
// 示例 2： 
//
// 输入：s = "aba"
//输出：2
//解释：一种最大拆分方法为 ['a', 'ba'] 。
// 
//
// 示例 3： 
//
// 输入：s = "aa"
//输出：1
//解释：无法进一步拆分字符串。
// 
//
// 
//
// 提示： 
//
// 
// 
// 1 <= s.length <= 16 
// 
// 
// s 仅包含小写英文字母 
// 
// 
// Related Topics 回溯算法 
// 👍 14 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxUniqueSplit(s string) int {

	ans := 0

	// 输入：s = "ababccc"
	//输出：5
	//解释：一种最大拆分方法为 ['a', 'b', 'ab', 'c', 'cc'] 。像 ['a', 'b', 'a', 'b', 'c', 'cc'] 这样
	//拆分不满足题目要求，因为其中的 'a' 和 'b' 都出现了不止一次。

	trace := func(rest string, cur []string) {}
	trace = func(rest string, cur []string) {
		if len(rest) == 0 && len(cur)>ans {
			ans = len(cur)
			return
		}

		for i:=1;i<=len(rest);i++ {
			elem := rest[:i]
			if !isOk(elem,cur) {
				continue
			}
			cur = append(cur,elem)
			rest = rest[i:]
			trace(rest,cur)
			cur = cur[:len(cur)-1]
			rest = elem + rest
		}
	}
	trace(s,[]string{})

	return ans
}

func isOk(elem string,cur []string) bool {
	for _,e :=range cur{
		if e == elem {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
