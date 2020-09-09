package main
//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。 
//
// 说明： 
//
// 
// 所有数字都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: k = 3, n = 7
//输出: [[1,2,4]]
// 
//
// 示例 2: 
//
// 输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]
// 
// Related Topics 数组 回溯算法 
// 👍 157 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	e := make([]int,0,9)
	for i := 1;i<10;i++{
		e = append(e,i)
	}
	trace := func(cur []int,sum,idx int) {}

	trace = func(cur []int,sum,idx int) {
		if sum>n || idx > len(e) || len(cur) > k {
			return
		}
		if sum == n && len(cur) == k {
			cp := make([]int,k)
			copy(cp,cur)
			ans = append(ans,cp)
			return
		}

		for i:=idx;i<len(e);i++ {
			cur = append(cur,e[i])
			trace(cur,sum+e[i],i+1)
			cur = cur[:len(cur)-1]
		}
	}
	trace([]int{},0,0)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
