package main
//Given an array of integers, return indices of the two numbers such that they a
//dd up to a specific target. 
//
// You may assume that each input would have exactly one solution, and you may n
//ot use the same element twice. 
//
// Example: 
//
// 
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
// 
// Related Topics 数组 哈希表 
// 👍 8849 👎 0


//leetcode submit region begin(Prohibit modification and deletion)


//存哈希表法
func twoSum(nums []int,target int) []int{
	m := make(map[int]int,len(nums))
	for index,iNum := range nums{
		m[iNum] = index
	}

	for jIndex,j:=range nums{
		if iIndex,ok:=m[target-j];ok&&(iIndex!=jIndex){
			return []int{jIndex,iIndex}
		}
	}
	return []int{0,0}
}

//暴力破解法
//func twoSum(nums []int, target int) []int {
//	for i,num:=range nums{
//		subTarget := target - num
//		for j:=i+1;j<len(nums);j++{
//			if subTarget == nums[j]{
//				return []int{i,j}
//			}
//		}
//	}
//	return []int{0,0}
//}
//leetcode submit region end(Prohibit modification and deletion)
