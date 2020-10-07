package main

import "sort"

//幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
//  输入： nums = [1,2,3]
// 输出：
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
// 
// Related Topics 位运算 数组 回溯算法 
// 👍 35 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	sort.Ints(nums)

	ans := [][]int{}

	trace := func(begin int,cur []int) {}
	trace = func(begin int,cur []int) {
		if begin > len(nums) {
			return
		}
		cp := make([]int,len(cur))
		copy(cp,cur)
		ans = append(ans,cp)

		for i:=begin;i<len(nums);i++ {
			cur = append(cur,nums[i])
			trace(i+1,cur)
			cur = cur[:len(cur)-1]
		}
	}
	trace(0,[]int{})

	return ans

}
//leetcode submit region end(Prohibit modification and deletion)
