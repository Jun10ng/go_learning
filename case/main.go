package main

import "fmt"

func main() {
	fmt.Println(StockCardExchange([]int{1, 2, 10}, 11))
	fmt.Println(StockCardExchange([]int{10}, 11))
	fmt.Println(StockCardExchange([]int{5, 6, 11}, 21))

}
func StockCardExchange(nums []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1 << 30
	}
	for i := 1; i <= amount; i++ {
		for _, num := range nums {
			if i-num < 0 {
				continue
			}
			dp[i] = min2(dp[i], dp[i-num]+1)
		}
	}
	if dp[amount] == 1<<30 {
		return -1
	}
	return dp[amount]
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
