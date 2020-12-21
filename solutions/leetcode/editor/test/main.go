package main

import (
	"fmt"
	"image"
)

func main() {
	fmt.Println(completeBag([]int{0, 3, 2, 1}, []int{0, 5, 2, 3}, 3, 5))
}

func completeBag(w, n []int, W, N int) int {
	/*
			tn
		   m 0 0 0 0
		     1
		     2
		 	 3
			 4
		     5

	*/
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dpi := make([]int, W+1)
		dp[i] = dpi
	}

	// 遍历每一件商品
	for i := 1; i < len(dp); i++ {
		// 遍历背包容量
		for rw := 1; rw < len(dp[i]); rw++ {
			//尽可能放入多次物品
		}
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
