package main

import (
	"fmt"
)

func main() {
	fmt.Println(productExceptSelf([]int{2, 1, 3, 4}))
}
func productExceptSelf(nums []int) []int {
	// 输入: [2,1,3,4]
	//输出: [12,24,8,6]
	/*
		pre[i] = 后 i 个元素的乘积，（i元素右边的乘积
		lp 前i-1个元素的乘积 (i元素左边的乘积
	*/
	ret := make([]int, 0)
	pre := make([]int, len(nums)+1)
	pre[0] = 1
	pre[len(nums)] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[len(nums)-i]
	}
	//1 4 12 12 1

	lp := 1
	for i := 0; i < len(nums); i++ {
		reti := lp * pre[len(nums)-(i+1)]
		ret = append(ret, reti)
		lp = lp * nums[i]
	}
	return ret
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
