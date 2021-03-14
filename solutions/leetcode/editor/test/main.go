package main

import (
	"fmt"
)

func main() {

	fmt.Print(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
func maxProfit(prices []int) int {
	/*
		0 -6 4 -2 3 -2
	*/
	ret := 0
	pre := []int{}
	pre = append(pre, 0)
	for i := 1; i < len(prices); i++ {
		pre = append(pre, prices[i]-prices[i-1])
	}
	for _, e := range pre {
		if e > 0 {
			ret += e
		}
	}
	return ret
}
func max2(a, b int) int {
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
