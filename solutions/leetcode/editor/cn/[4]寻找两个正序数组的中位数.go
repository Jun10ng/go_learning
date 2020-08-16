package main

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//
// 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。 
//
// 你可以假设 nums1 和 nums2 不会同时为空。 
//
// 
//
// 示例 1: 
//
// nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
// 
//
// 示例 2: 
//
// nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5
//
// [1,3,5] [2,4,6] (3+4)/2 3.5
// [1,3,5] [2,4]   3
//
// Related Topics 数组 二分查找 分治算法 
// 👍 3063 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var ret float64

	l1 ,l2:= len(nums1), len(nums2)
	lenSum := l1 + l2
	mid := lenSum/2
	i1,i2 := 0,0

	for ;(i1+i2+2)<=mid; {
		if nums1[i1]<=nums2[i2] {
			i1++
		}else {
			i2++
		}
	}
	if lenSum%2 ==0{
		ret = (float64(nums1[i1])+float64(nums2[i2]))/2
	}else {
		if nums1[i1]>nums2[i2] {
			ret = float64(nums2[i2])
		}else {
			ret = float64(nums1[i1])
		}
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
