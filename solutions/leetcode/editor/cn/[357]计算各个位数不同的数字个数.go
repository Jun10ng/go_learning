package main
//给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n 。 
//
// 示例: 
//
// 输入: 2
//输出: 91 
//解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。
// 
// Related Topics 数学 动态规划 回溯算法 
// 👍 90 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func countNumbersWithUniqueDigits(n int) int {
	cnt := 0
	cur := make([]int,0)
	/*
		100
		      0 ...
		     0    1 ....
		    x 1  0 1 2...

	*/
	trace := func(cur []int) {}

	trace = func(cur []int) {
		if len(cur) == n{
			cnt++
			return
		}

		if len(cur) > n {
			return
		}

		for i:=0;i<=9;i++ {
			if isOK(cur,i) {
				cur = append(cur,i)
				trace(cur)
				cur = cur[:len(cur)-1]
			}
		}
	}
	trace(cur)
	return cnt
}

// true => 没包含
func isOK(cur []int,e int) bool {
	zeroFlag :=true
	for _,c := range cur{
		if c==0 && zeroFlag{
			continue
		}
		if c!=0 {
			zeroFlag = false
		}
		if c == e && !zeroFlag{
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
