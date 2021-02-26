package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
func pivotIndex(nums []int) int {
	/*
		[1, 7, 3, 6, 5, 6] 3
		sub[i] i元素后面的和
		leftSum 前i个元素的和
	*/
	sub := make([]int, len(nums)+1)
	sub[len(nums)] = 0
	for i := 1; i < len(nums); i++ {
		num := nums[len(nums)-i]
		sub[len(nums)-i] = sub[len(nums)-i+1] + num
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if leftSum == sub[i+1] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	tmp := min2(a, b)
	return min2(tmp, c)
}
