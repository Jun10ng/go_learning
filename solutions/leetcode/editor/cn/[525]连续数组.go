package main
//给定一个二进制数组, 找到含有相同数量的 0 和 1 的最长连续子数组（的长度）。 
//
// 
//
// 示例 1: 
//
// 输入: [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量0和1的最长连续子数组。 
//
// 示例 2: 
//
// 输入: [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。 
//
// 
//
// 注意: 给定的二进制数组的长度不会超过50000。 
// Related Topics 哈希表 
// 👍 236 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength(nums []int) int {
	indices := map[int]int{0: -1}
	var res, prefixSum int
	for i, v := range nums {
		if v == 0 {
			v = -1
		}
		prefixSum += v
		if j, ok := indices[prefixSum]; !ok {
			indices[prefixSum] = i
		} else if i-j > res {
			res = i - j
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
