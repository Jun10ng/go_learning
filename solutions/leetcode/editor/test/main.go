package main

import "fmt"

func main() {
	m := []int{1, 2} // 1357 1347
	//m := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1} // 1357 1347
	fmt.Println(jump(m))
}
func jump(nums []int) int {
	/*
		{2,3,1,1,4} => 2
		dp[i] 跳到 i 的最少次数
		0  1 1 2 2
	*/
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 99999
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := 1; j < num+1; j++ {
			if i+j >= len(nums) {
				continue
			}
			dp[i+j] = min2(dp[i+j], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
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
