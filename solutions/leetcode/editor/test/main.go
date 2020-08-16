package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var ret float64
	var isOdd bool
	var i, j int
	mid := (len(nums1) + len(nums2) - 2) / 2
	if (len(nums1)+len(nums2))%2 == 0 {
		isOdd = false
	} else {
		isOdd = true
	}
	//i,j = mid/2,mid/2

	for mid >= 1 {
		i, j = mid/2, mid/2
		if nums1[i] < nums2[j] {
			mid = mid - i - 1
			nums1 = nums1[mid+1:]
		} else {
			mid = mid - j - 1
			nums2 = nums2[mid+1:]
		}
	}

	if !isOdd {
		//偶数
		ret = (float64(nums1[0]) + float64(nums2[0])) / 2
	} else {
		if nums1[0] > nums2[0] {
			ret = float64(nums2[0])
		} else {
			ret = float64(nums1[0])
		}
	}
	return ret
}
func main() {
	n := []int{-1, 3}
	m := []int{1, 2}
	fmt.Print(findMedianSortedArrays(n, m))
}
