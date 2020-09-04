package main

import "sort"

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。 
//
// 说明： 
//
// 
// 所有数字（包括 target）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1： 
//
// 输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
// 
//
// 示例 2： 
//
// 输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//] 
//
// 
//
// 提示： 
//
// 
// 1 <= candidates.length <= 30 
// 1 <= candidates[i] <= 200 
// candidate 中的每个元素都是独一无二的。 
// 1 <= target <= 500 
// 
// Related Topics 数组 回溯算法 
// 👍 850 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func combinationSum(candidates []int, target int) [][]int {
	ans := &[][]int{}
	e := &[]int{}
	cur := 0
	sort.Ints(candidates)
	trace(candidates,ans,e,cur,target,0)
	return *ans
}

func trace(c []int,ans *[][]int, e *[]int,cur,t,begin int)  {
	if cur == t{
		cp := make([]int,len(*e))
		copy(cp,*e)
		*ans = append(*ans,cp)
		return
	}

	dlt := t - cur
	for i:=begin;i<len(c);i++ {
		if c[i]<=dlt {
			*e = append(*e,c[i])
			trace(c,ans,e,cur+c[i],t,i)
			*e = (*e)[:len(*e)-1]
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
