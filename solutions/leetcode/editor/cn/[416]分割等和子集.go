package main
//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。 
//
// 注意: 
//
// 
// 每个数组中的元素不会超过 100 
// 数组的大小不会超过 200 
// 
//
// 示例 1: 
//
// 输入: [1, 5, 11, 5]
//
//输出: true
//
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
// 
//
// 
//
// 示例 2: 
//
// 输入: [1, 2, 3, 5]
//
//输出: false
//
//解释: 数组不能分割成两个元素和相等的子集.
// 
//
// 
// Related Topics 动态规划 
// 👍 622 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	/*
			特殊 0-1 背包问题 价值都为1，即相等价值
			[1, 5, 11, 5] true
			[1,2,5] false

		    状态 n:背包内物品数量 w:背包还能装下的重量
			tn: 前n件商品
			dp[tn][w] : 在选择商品范围为前n件商品时，背包有w空间，此时最大价值
		    tn
		  w 0 0 0 0 0
			0 1 1 1 1
		    0 1 2 3 3
		    0 1 2 3 3

		   return dp[len(nums)][sum] == sum
	*/

	//初始化
	sum := 0
	for _,e:=range nums{
		sum += e
	}
	if sum&1!=0 {
		return false
	}
	sum = sum/2
	dp := make([][]int,0)
	for i:=0;i<=len(nums);i++ {
		dpi := make([]int,sum+1)
		dp = append(dp,dpi)
	}
	for i:=0;i<=len(nums);i++ {
		dp[i][0]=0
	}
	for i:=0;i<=sum;i++ {
		dp[0][i]=0
	}

	// 决策
	for tn:=1;tn<=len(nums);tn++{
		curNum := nums[tn-1]
		for w:=1;w<=sum;w++ {
			if curNum>w {
				continue
			}
			dp[tn][w] = max(dp[tn-1][w],dp[tn-1][w-curNum]+curNum)
		}
	}

	return dp[len(nums)][sum]==sum
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
