package main

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 3000 
// -105 <= nums[i] <= 105 
// 
// Related Topics 数组 双指针 
// 👍 3189 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int,0)
	//m := make(map[int]int,len(nums))
	//for index:=0;index<len(nums);index++{
	//	value := nums[index]
	//	m[value] = index
	//}

	for i:=0;i<len(nums);i++ {
		if nums[i]>0 {
			return ret
		}
		if i>0 && nums[i]==nums[i-1] {
			continue
		}
		l,r :=i+1,len(nums)-1
		for l<r {
			s := nums[i]+nums[l]+nums[r]
			if s<0{
				l++
				for l<r && nums[l]==nums[l-1]{l++}
			}else if s>0 {
				r--
				for l<r && nums[r]==nums[r+1]{r--}
			}else {
				ret = append(ret,[]int{nums[i],nums[l],nums[r]})
				l++
				r--
				for l<r && nums[l]==nums[l-1]{l++}
				for l<r && nums[r]==nums[r+1]{r--}
			}
		}

	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
