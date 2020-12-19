package main
//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。 
//
// 
// 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。 
//
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
//输出：4
//解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
//其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 
//n 的值 3 。
// 
//
// 示例 2： 
//
// 
//输入：strs = ["10", "0", "1"], m = 1, n = 1
//输出：2
//解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= strs.length <= 600 
// 1 <= strs[i].length <= 100 
// strs[i] 仅由 '0' 和 '1' 组成 
// 1 <= m, n <= 100 
// 
// Related Topics 动态规划 
// 👍 294 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findMaxForm(strs []string, m int, n int) int {
	/*
			0-1 背包问题 找出最大子集

			状态：m 和 n 和 str[i] 其实是有三种状态的，
				但由于使用了状态压缩，去掉了str[i]
			dp[m][n] : 满足条件m n的最大子集
			dp[0][0] = 0
			k= str[i]含0数 z=str[i]含1数
		    dp[m][n] = max{dp[m][n],dp[m-k][n-z]+1}

		难点：由于状态压缩，要对m和n，从大到小遍历
		在状态压缩时，为了避免后面的值被反复更新，要时候倒序遍历
	*/
	dp := make([][]int,0)
	for i:=0;i<=m;i++{
		dpi := make([]int,n+1)
		dp = append(dp,dpi)
	}

	for _,str:=range strs{
		k,z := gen01(str)
		for i:=m;i>=k;i-- {
			for j:=n;j>=z;j-- {
				dp[i][j] = max(dp[i][j],dp[i-k][j-z]+1)
			}
		}
	}
	return dp[m][n]
}
func gen01(s string) (int,int) {
	k,z:=0,0
	for _,e:=range s{
		if string(e)=="0" {
			k++
		}else {
			z++
		}
	}
	return k,z
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
