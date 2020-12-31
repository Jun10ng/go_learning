package main
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。 
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。 
//
// 问总共有多少条不同的路径？ 
//
// 
//
// 示例 1： 
//
// 
//输入：m = 3, n = 7
//输出：28 
//
// 示例 2： 
//
// 
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
// 
//
// 示例 3： 
//
// 
//输入：m = 7, n = 3
//输出：28
// 
//
// 示例 4： 
//
// 
//输入：m = 3, n = 3
//输出：6 
//
// 
//
// 提示： 
//
// 
// 1 <= m, n <= 100 
// 题目数据保证答案小于等于 2 * 109 
// 
// Related Topics 数组 动态规划 
// 👍 842 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	// dp[i][j] 从起点 到达i，j 有几种走法
	dp := make([][]int,0)
	for i := 0; i < m+1; i++ {
		dpi := make([]int,n+1)
		dp = append(dp,dpi)
	}
	dp[0][1] = 1
	for i := 1; i < m+1	; i++ {
		for j:=1;j<n+1;j++ {
			dp[i][j] += dp[i-1][j]+dp[i][j-1]
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
