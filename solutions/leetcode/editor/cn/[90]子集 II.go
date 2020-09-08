package main

import "sort"

//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
// 输入: [1,2,2]
//输出:
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//] 
// Related Topics 数组 回溯算法 
// 👍 299 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	trace := func(cur []int, begin int) {}

	trace = func(cur []int, begin int) {
		if begin > len(nums) {
			return
		}
		cp := make([]int,len(cur))
		copy(cp,cur)
		ans = append(ans,cp)

		for i:=begin;i<len(nums);i++ {
			if i > begin && nums[i]==nums[i-1] {
				continue
			}
			cur = append(cur,nums[i])
			trace(cur,i+1)
			cur = cur[:len(cur)-1]
		}
	}

	trace([]int{},0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
