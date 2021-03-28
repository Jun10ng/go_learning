package main

import (
	"reflect"
	"sort"
)

//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//d
// 示例: 
//
// 输入: [1,1,2]
//输出:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//] 
// Related Topics 回溯算法 
// 👍 381 👎 0


//leetcode submit region begin(Prohibit modification and deletion)


func permuteUnique(nums []int) [][]int {
	//var ans [][]int
	sort.Ints(nums)
	ans := make([][]int,0,len(nums)*len(nums))
	t := make([]int,0,len(nums))
	used := make([]bool,len(nums),len(nums))
	trace(nums,t,used,&ans)
	return ans
}
func trace(nums []int, t []int,used []bool,ans *[][]int) {
	if len(nums) == len(t) {
		ct := make([]int,len(t),len(t))
		copy(ct,t)
		*ans = append(*ans, ct)
		return
	}
	for i,e := range nums{
		if (used[i]) {
			continue
		}
		if i!=0 && used[i-1] && e == nums[i-1]{
			continue
		}
		used[i] = true
		t = append(t,e)
		trace(nums,t,used,ans)
		used[i] = false
		t = removeLast(t)
	}
	return
}

func removeLast(t []int) []int {
	return t[:len(t)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
