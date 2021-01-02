package main

import "fmt"

func main() {
	m := []int{1, 3, 5, 4, 7} // 1357 1347
	fmt.Println(findNumberOfLIS(m))

	//m1 := []int{1,2,4,3,5,4,7,2} // 12457 12456 12357
	//fmt.Println(findNumberOfLIS(m1))

	//m2 := []int{2, 15, 3, 7, 8, 6, 18} //  2 3 7 8 18
	//
	//fmt.Println(lengthOfLIS(m1))
	//fmt.Println(lengthOfLIS(m2))
	//fmt.Println(canJump([]int{3,2,1,0,4}))
}
func findNumberOfLIS(nums []int) int {
	/*
		dp[i] 从起始位置到i元素的最长上升序列
		count[i] 从起始位置到i元素的最长上升序列个数
	*/
	numOfLIS, maxLen := 0, 1
	dp := make([]int, len(nums))
	// 表示以第 i 个数字结尾的序列的最长上升子序列的数量。
	// 解题思路：相当于对count 做一次DP
	count := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
		count[i] = 1
	}

	for j := 0; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				switch {
				case dp[i]+1 > dp[j]:
					// 此时需要更新dp
					dp[j] = dp[i] + 1
					count[j] = count[i] // 最长子序列在延续
				case dp[i]+1 == dp[j]:
					count[j] += count[i] //出现了长度相同的最长子序列，叠加
				}
			}
		}
		maxLen = max(dp[j], maxLen)
	}
	for i, e := range dp {
		if e == maxLen {
			numOfLIS += count[i]
		}
	}
	return numOfLIS
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
