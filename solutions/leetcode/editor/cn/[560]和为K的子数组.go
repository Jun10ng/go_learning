package main
//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。 
//
// 示例 1 : 
//
// 
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
// 
//
// 说明 : 
//
// 
// 数组的长度为 [1, 20,000]。 
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。 
// 
// Related Topics 数组 哈希表 
// 👍 778 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	/*
		前缀和
		pre[i] 代表前i个元素之和
		num = nums[i]
		ret = num + pre[
	*/
	ret := 0
	sum := 0
	pre := make(map[int]int,0)
	pre[0] = 1
	for i := 0;i<len(nums);i++ {
		sum += nums[i]
		if _,ok := pre[sum-k];ok {
			ret+=pre[sum-k]
		}
		pre[sum]+=1
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
