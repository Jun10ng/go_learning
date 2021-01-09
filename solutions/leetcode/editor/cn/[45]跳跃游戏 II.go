package main
//给定一个非负整数数组，你最初位于数组的第一个位置。 
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。 
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。 
//
// 示例: 
//
// 输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 
//
// 说明: 
//
// 假设你总是可以到达数组的最后一个位置。 
// Related Topics 贪心算法 数组 
// 👍 789 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	/*
		{2,3,1,1,4} => 2
		dp[i] 跳到 i 的最少次数
		0  1 1 2 2
	*/
	if len(nums)==1 {
		return 0
	}
	dp := make([]int,len(nums))
	for i,_:=range dp{
		dp[i]= 99999
	}
	dp[0] = 0
	for i := 0; i <len(nums) ; i++ {
		num := nums[i]
		for j :=1;j<num+1;j++ {
			if i+j >= len(nums) {
				continue
			}
			dp[i+j] = min2(dp[i+j],dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
