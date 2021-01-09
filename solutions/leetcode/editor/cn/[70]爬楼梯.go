package main
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。 
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？ 
//
// 注意：给定 n 是一个正整数。 
//
// 示例 1： 
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶 
//
// 示例 2： 
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
// 
// Related Topics 动态规划 
// 👍 1403 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	/*
			dp[i] 到达i位置有几种跳法
		    决策 使用1 跳 还是2 跳
		    1 | 1,2,1
	*/

	dp := make([]int,n+1)
	dp[0]= 1
	stepList := []int{1,2}
	for i:=1;i<n+1;i++ {
		for _,step := range stepList{
			if i-step >=0  {
				dp[i] += dp[i-step]
			}
		}
	}
	return dp[n]
}
//leetcode submit region end(Prohibit modification and deletion)
