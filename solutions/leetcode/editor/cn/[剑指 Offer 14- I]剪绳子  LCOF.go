package main

//English description is not available for the problem. Please switch to Chinese
//. Related Topics 数学 动态规划 
// 👍 91 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
	if n==1 ||n ==2 {
		return 1
	}
	dp := make([]int,n+1)
	dp[1],dp[2] = 1,1
	for i :=3;i<=n;i++ {
		for j:=1;j<i;j++{
			dp[i] = Max(dp[i],Max(j,dp[j])*Max(i-j,dp[i-j]))
		}
	}
	return dp[n]

}
func Max(a,b int) int {
	if a<=b{
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
