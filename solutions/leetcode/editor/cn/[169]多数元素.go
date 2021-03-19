package main
//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。 
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：[3,2,3]
//输出：3 
//
// 示例 2： 
//
// 
//输入：[2,2,1,1,1,2,2]
//输出：2
// 
//
// 
//
// 进阶： 
//
// 
// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。 
// 
// Related Topics 位运算 数组 分治算法 
// 👍 917 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	/*
		{2,2,1,1,1,2,2}
		[2 2 1] [1 1 2 2]

	*/
	if len(nums)==1 {
		return nums[0]
	}
	m := len(nums)/2
	l := majorityElement(nums[:m])
	r := majorityElement(nums[m:])
	if l == r {
		return l
	}
	lc,rc := 0,0
	for _,e:=range nums{
		if e==l {
			lc++
		}else {
			rc++
		}
	}
	if lc>rc{
		return l
	}
	return r
}
//leetcode submit region end(Prohibit modification and deletion)
