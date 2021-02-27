package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(minSumOfLengths([]int{1,2,2,3,2,6,7,2,1,4,},5))
	//fmt.Println(minSumOfLengths([]int{1,6,1},7))
	//fmt.Println(minSumOfLengths([]int{3,1,1,1,5,1,2,1},3))
	fmt.Println(minSumOfLengths([]int{1, 1, 1, 2, 2, 2, 4, 4}, 6))
}
func minSumOfLengths(arr []int, target int) int {
	/*
		参考：https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/solution/xiang-xi-jiang-jie-yi-xia-shuang-zhi-zhe-jjt9/
		双指针 动态规划
		通过双指针来找符合条件的子数组
		通过动态规划来判断是否越界
		dp[i]: i元素之前满足的最小子数组长度
	*/
	dp := make([]int, len(arr)+1)
	dp[0] = len(arr) + 1
	allAns := make([]int, 0)
	i := 0
	j := 0
	sum := 0
	for ; j < len(arr); j++ {
		sum += arr[j]
		for sum > target {
			sum -= arr[i]
			i++
		}
		if sum == target {
			length := j - i + 1
			allAns = append(allAns, length+dp[i])
			dp[j+1] = min2(dp[j], length)
		} else {
			dp[j+1] = dp[j]
		}
	}
	sort.Ints(allAns)
	if len(allAns) < 2 || allAns[0] > len(arr) {
		return -1
	}
	return allAns[0]
}
func sumOfSlice(s []int, i, j int) int {
	sum := s[j]
	for _, e := range s[i:j] {
		sum += e
	}
	return sum
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
