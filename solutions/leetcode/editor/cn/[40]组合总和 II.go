package main

import "sort"

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。 
//
// 说明： 
//
// 
// 所有数字（包括目标数）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
// 
//
// 示例 2: 
//
// 输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//] 
// Related Topics 数组 回溯算法 
// 👍 349 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	ans := &[][]int{}
	e := &[]int{}
	cur := 0
	sort.Ints(candidates)
	//used := make([]bool,len(candidates))
	trace(candidates,ans,e,cur,target,0)
	return *ans
}

func trace(c []int,ans *[][]int,e *[]int,cur,t,begin int)  {
	if cur == t{
		cp := make([]int,len(*e))
		copy(cp,*e)
		*ans = append(*ans,cp)
		return
	}
	if cur > t{
		return
	}
	for i:=begin;i<len(c);i++ {
		if i>begin&&c[i]==c[i-1] {
			continue
		}
		//(*used)[i] = true
		*e = append(*e,c[i])
		trace(c,ans,e,cur+c[i],t,i+1)
		*e = (*e)[:len(*e)-1]
		//(*used)[i] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
