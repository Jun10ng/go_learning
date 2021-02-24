package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}
func subarraySum(nums []int, k int) int {
	/*
		前缀和
		pre[i] 代表前i个元素之和
		num = nums[i]
		ret = num + pre[
	*/
	ret := 0
	sum := 0
	pre := make(map[int]int, 0)
	pre[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := pre[sum-k]; ok {
			ret += pre[sum-k]
		}
		pre[sum] += 1
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
