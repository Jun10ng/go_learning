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
// [1,3,5,8] [2,4,6,7] (5+4)/2 4.5
// [1,3,5] [2,4]   3
//
// Related Topics 数组 二分查找 分治算法 
// 👍 3063 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 这道题的实质是取两个数组的第K个最小值
// 长度为偶数时，就是第K小和第K+1小值 相加除二
// 如果使用双指针的解法没法解决两个中位数都在一个数组内
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	findKth := func(k int)float64{
		i,j := 0,0
		for {
			//某个序列为空
			if 0 == len(nums1){
				return float64(nums2[j + k -1])
			}
			if 0 == len(nums2){
				return float64(nums1[i + k -1])
			}
			if k==1{
				return float64(min2(nums1[i],nums2[j]))
			}
			halfK := k/2
			newi := min2(i + halfK+1,len(nums1)-1)
			newj := min2(j + halfK+1,len(nums2)-1)
			if nums1[newi]<nums2[newj]{
				k = k- (newi - i + 1)
				i = newi + 1
			}else {
				k = k- (newj - j + 1)
				j = newj + 1
			}
		}
	}
	l1,l2 := len(nums1),len(nums2)
	isEven := (l1+l2) & 1 == 0
	mid := (l1+l2)/2
	if isEven{
		return (findKth(mid+1)+findKth(mid))/2
	}
	return findKth(mid+1)
}

func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
