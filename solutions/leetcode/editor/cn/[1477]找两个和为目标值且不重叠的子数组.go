package main

import "sort"

//给你一个整数数组 arr 和一个整数值 target 。
//
// 请你在 arr 中找 两个互不重叠的子数组 且它们的和都等于 target 。可能会有多种方案，请你返回满足要求的两个子数组长度和的 最小值 。 
//
// 请返回满足要求的最小长度和，如果无法找到这样的两个子数组，请返回 -1 。 
//
// 
//
// 示例 1： 
//
// 输入：arr = [3,2,2,4,3], target = 3
//输出：2
//解释：只有两个子数组和为 3 （[3] 和 [3]）。它们的长度和为 2 。
// 
//
// 示例 2： 
//
// 输入：arr = [7,3,4,7], target = 7
//输出：2
//解释：尽管我们有 3 个互不重叠的子数组和为 7 （[7], [3,4] 和 [7]），但我们会选择第一个和第三个子数组，因为它们的长度和 2 是最小值。
// 
//
// 示例 3： 
//
// 输入：arr = [4,3,2,6,2,3,4], target = 6
//输出：-1
//解释：我们只有一个和为 6 的子数组。
// 
//
// 示例 4： 
//
// 输入：arr = [5,5,4,4,5], target = 3
//输出：-1
//解释：我们无法找到和为 3 的子数组。
// 
//
// 示例 5： 
//
// 输入：arr = [3,1,1,1,5,1,2,1], target = 3
//输出：3
//解释：注意子数组 [1,2] 和 [2,1] 不能成为一个方案因为它们重叠了。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr.length <= 10^5 
// 1 <= arr[i] <= 1000 
// 1 <= target <= 10^8 
// 
// Related Topics 动态规划 
// 👍 56 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minSumOfLengths(arr []int, target int) int {
	/*
		参考：https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/solution/xiang-xi-jiang-jie-yi-xia-shuang-zhi-zhe-jjt9/
		双指针 动态规划
		通过双指针来找符合条件的子数组
		通过动态规划来判断是否越界
		dp[i]: i元素之前满足的最小子数组长度
	*/
	dp :=  make([]int,len(arr)+1)
	dp[0] = len(arr)+1
	allAns := make([]int,0)
	i:=0
	j:=0
	sum := 0
	for ;j<len(arr);j++{
		sum += arr[j]
		for sum>target {
			sum -= arr[i]
			i++
		}
		if sum == target {
			length := j-i+1
			allAns = append(allAns,length+dp[i])
			dp[j+1] = min2(dp[j],length)
		}else {
			dp[j+1] = dp[j]
		}
	}
	sort.Ints(allAns)
	if len(allAns)<2||allAns[0]>len(arr) {
		return -1
	}
	return allAns[0]
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}


//leetcode submit region end(Prohibit modification and deletion)
