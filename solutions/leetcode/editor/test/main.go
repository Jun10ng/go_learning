package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1}, 1))
}

func findTargetSumWays(nums []int, S int) int {
	/*
			状态压缩
		https://leetcode-cn.com/problems/target-sum/solution/huan-yi-xia-jiao-du-ke-yi-zhuan-huan-wei-dian-xing/
		把所有符号为正的数总和设为一个背包的容量，容量为x；把所有符号为负的数总和设为一个背包的容量，
		容量为y。在给定的数组中，有多少种选择方法让背包装满。令sum为数组的总和，
		则x+y=sum。而两个背包的差为S,
		则x-y=S。从而求得x=(S+sum)/2。
		基于上述分析，题目转换为背包问题：给定一个数组和一个容量为x的背包，求有多少种方式让背包装满。
	*/

	sum := 0
	for _, e := range nums {
		sum += e
	}
	if (S+sum)&1 != 0 || S > sum {
		return 0
	}

	length := (S + sum) / 2
	dp := make([]int, 0)
	for i := 0; i <= length; i++ {
		dp = append(dp, 0)
	}
	dp[0] = 1

	for _, num := range nums {
		for i := length; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[length]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
