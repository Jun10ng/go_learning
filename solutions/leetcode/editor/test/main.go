package main

import (
	"fmt"
)

func main() {
	//m := [][]int{
	//	{0, 0}, {1, 1}, {0, 0},
	//}
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	//fmt.Println(canJump([]int{3,2,1,0,4}))
}
func maxProfit(prices []int) int {
	//   [3,3,5,0,0,3,1,4] => 6
	/*
			状态：天数，是否持有股票，交易次数
			dp[i][j][k] : 在第i天交易某个股票
		    j = 0 不持有 j = 1持有
			dp[i][j][k]

			dp[i][0][0] = 0                   , j=0,k=0
		    //卖出过一次，最大利润有两种可能，要么是当天卖出，要么是之前卖出
		    dp[i][0][1]= max{ dp[i-1][1][0]+p[i], dp[i-1][0][1] }
		    dp[i][0][2]= max{ dp[i-1][1][1]+p[i], dp[i-1][0][2] }
		    // 买入过一次 股票是当天买入，or 股票之前就有
		    dp[i][1][0]= max{ dp[i-1][0][0]-p[i], dp[i-1][1][0] }
		   	dp[i][1][1]= max{ dp[i-1][0][1]-p[i], dp[i-1][1][1] }
			dp[i][1][2] = -99999

	*/

	dp := make([][][]int, 0)
	for i := 0; i < len(prices); i++ {
		dpi := make([][]int, 0)
		for j := 0; j < 2; j++ {
			dpj := make([]int, 3)
			dpi = append(dpi, dpj)
		}
		dp = append(dp, dpi)
	}

	// 初始化
	// 处理第一天
	//假设第一天没有买入
	//dp[0][0][0] = 0;
	//dp[0][0][1] = 0;
	//dp[0][0][2] = 0;

	//第一天的k=1 k=2的情况都要处理下
	dp[0][1][0] = -prices[0]
	dp[0][1][1] = -prices[0]
	dp[0][0][2] = -prices[0]

	for i := 1; i < len(prices); i++ {
		p := prices[i]
		dp[i][0][0] = 0
		//卖出过一次，最大利润有两种可能，要么是当天卖出，要么是之前卖出
		dp[i][0][1] = max(dp[i-1][1][0]+p, dp[i-1][0][1])
		dp[i][0][2] = max(dp[i-1][1][1]+p, dp[i-1][0][2])
		// 买入过一次 股票是当天买入，or 股票之前就有
		dp[i][1][0] = max(dp[i-1][0][0]-p, dp[i-1][1][0])
		dp[i][1][1] = max(dp[i-1][0][1]-p, dp[i-1][1][1])
		dp[i][1][2] = 0
	}
	return max(dp[len(prices)-1][0][1], dp[len(prices)-1][0][2])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
