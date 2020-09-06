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
	ans := &[][]int{}
	e := &[]int{}

	trace(ans,e,nums,0)
	return *ans
}

func trace(ans *[][]int,e *[]int, nums []int, begin int)  {
	if len(*e)<= len(nums){
		cp := make([]int,len(*e))
		copy(cp,*e)
		*ans = append(*ans,cp)
	}
	for i:=begin;i<len(nums);i++ {
		*e = append(*e,nums[i])
		trace(ans,e,nums,i+1)
		*e = (*e)[:len(*e)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
