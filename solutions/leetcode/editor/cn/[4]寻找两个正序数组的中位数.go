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
//这道题的实质是取两个数组的第K个最小值
//
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sumLen := len(nums2)+len(nums1)
	k := sumLen/2
	if sumLen % 2 == 0{
		return float64((getKthMin(nums1,nums2,k)+getKthMin(nums1,nums2,k+1)))/2.0
	}else{
		return float64(getKthMin(nums1,nums2,k+1))
	}
}

func getKthMin(num1,num2 []int,k int) int {
	i,j := 0,0
	for{
		//某个序列为空
		if i == len(num1){
			return num2[j + k -1]
		}
		if j == len(num2){
			return num1[i + k -1]
		}
		// 出口
		if k == 1 {
			return min(num1[i], num2[j])
		}
		//正常情况
		half := k/2
		newi := min(i+half,len(num1))-1
		newj := min(j+half,len(num2))-1
		p1,p2 := num1[newi],num2[newj]
		if p1<=p2{
			k = k - (newi - i + 1)
			i = newi + 1
		}else {
			k = k - (newj - j + 1)
			j = newj + 1
		}
	}
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
