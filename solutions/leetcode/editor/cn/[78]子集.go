package main
//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。 
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
// 输入: nums = [1,2,3]
//输出:
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
// Related Topics 位运算 数组 回溯算法 
// 👍 738 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func subsets(nums []int) [][]int {
	ret := make([][]int,0)
	ln := len(nums)
	vis := make([]bool,len(nums))
	doSubSet := func(set []int,begin int) {}
	doSubSet = func(set []int,begin int) {
		if len(set) <= ln{
			cpy:=make([]int,len(set))
			copy(cpy,set)
			ret = append(ret,cpy)
		}
		if len(set) == ln {
			return
		}

		for i:=begin;i<len(nums);i++{
			if !vis[i]{
				vis[i] = true
				doSubSet(append(set,nums[i]),i+1)
				vis[i] = false
			}
		}
	}
	doSubSet([]int{},0)
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
