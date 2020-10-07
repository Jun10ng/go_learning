package main
//括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。 
//
// 说明：解集不能包含重复的子集。 
//
// 例如，给出 n = 3，生成结果为： 
//
// 
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
// 
// Related Topics 字符串 回溯算法 
// 👍 35 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var maxLeft int
func generateParenthesis(n int) []string {
	ans := make([]string,0)
	maxLeft = n
	//leftCnt := 3
	trace := func(leftCnt int,cur string) {}
	trace = func(leftCnt int, cur string) {
		if len(cur)==2 * n {
			cp := cur
			ans = append(ans,cp)
			return
		}
		// "("
		if isOk(leftCnt+1,cur+"(") {
			trace(leftCnt+1,cur+"(")
		}

		//")"
		if isOk(leftCnt,cur+")") {
			trace(leftCnt,cur+")")
		}
	}
	trace(0,"")
	return ans
}
func isOk(leftCnt int,cur string) bool {
	if len(cur)>2 * maxLeft {
		return false
	}
	if leftCnt > maxLeft{
		return false
	}
	if len(cur)-leftCnt > maxLeft {
		return false
	}
	if len(cur) == 2 * maxLeft {
		//判断长度为2n的字符串是否合格
		//()(())
		leftStack := ""
		for _,e := range cur{
			if string(e) == "(" {
				leftStack = leftStack + string(e)
			}else {
				if len(leftStack)==0 {
					return false
				}
				leftStack = leftStack[:len(leftStack)-1]
			}
		}
		return len(leftStack) == 0
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
