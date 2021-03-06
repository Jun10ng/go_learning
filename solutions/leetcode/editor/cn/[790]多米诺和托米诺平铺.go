package main
//有两种形状的瓷砖：一种是 2x1 的多米诺形，另一种是形如 "L" 的托米诺形。两种形状都可以旋转。 
//
// 
//XX  <- 多米诺
//
//XX  <- "L" 托米诺
//X
// 
//
// 给定 N 的值，有多少种方法可以平铺 2 x N 的面板？返回值 mod 10^9 + 7。 
//
// （平铺指的是每个正方形都必须有瓷砖覆盖。两个平铺不同，当且仅当面板上有四个方向上的相邻单元中的两个，使得恰好有一个平铺有一个瓷砖占据两个正方形。） 
//
// 
//示例:
//输入: 3
//输出: 5
//解释: 
//下面列出了五种不同的方法，不同字母代表不同瓷砖：
//XYZ XXZ XYY XXY XYY
//XYZ YYZ XZZ XYY XXY 
//
// 提示： 
//
// 
// N 的范围是 [1, 1000] 
// 
//
// 
// Related Topics 动态规划 
// 👍 76 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func numTilings(N int) int {
	/*
		输入 3 输出  5

		dp[i] 为 搭建方案种数
		dp[0] = 0 dp[1] =1 dp[2]=2 dp[3] = 5

		dp[i] = max(dp[i-j] * dp[j]) 0<j<i
	*/

	dp := make([]int,N+1)
	dp[0],dp[1],dp[2],dp[3] = 0,1,2,5

	for i:=4;i<N+1;i++ {
		for j:=1;j<i;j++ {
			dp[i] = max(dp[i],dp[i-j]*dp[j])
		}
	}

	return dp[N]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
