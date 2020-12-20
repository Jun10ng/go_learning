package main
//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选
//择一个符号添加在前面。 
//
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。 
//
// 
//
// 示例： 
//
// 输入：nums: [1, 1, 1, 1, 1], S: 3
//输出：5
//解释：
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//
//一共有5种方法让最终目标和为3。
// 
//
// 
//
// 提示： 
//
// 
// 数组非空，且长度不会超过 20 。 
// 初始的数组的和不会超过 1000 。 
// 保证返回的最终结果能被 32 位整数存下。 
// 
// Related Topics 深度优先搜索 动态规划 
// 👍 496 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, S int) int {
	/*
	 所有正符号的数之和 x
	 所有负符号的数之和 y
	x+y = sum
	x -y = S
	所以 x = (sum+S)/2
	问题转换为，一个容量为X的背包，能有几种方法放满
	*/
	sum := 0
	for _,e:=range nums{
		sum += e
	}
	if (sum+S)&1!=0 || sum < S {
		return 0
	}
	x := (sum+S)/2
	dp := make([]int,x+1)
	dp[0] = 1 // 0容量 1 种办法 初始化

	for _,num:=range nums{
		for i := x; i >= num; i-- {
			// 状态压缩 所以要从大到小遍历
			dp[i] += dp[i-num]
		}
	}
	return dp[x]

}
//leetcode submit region end(Prohibit modification and deletion)
