package main
//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
//例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
//
// 示例 1： 
//
// 输入：[3,4,5,1,2]
// 输出：1
// 
//
// 示例 2： 
//
// 输入：[2,2,2,0,1]
// 输出：0
// 
//
// 注意：本题与主站 154 题相同：https://leetcode-cn.com/problems/find-minimum-in-rotated-sor
//ted-array-ii/ 
// Related Topics 二分查找 
// 👍 135 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minArray(numbers []int) int {
	if len(numbers)==1{
		return numbers[0]
	}
	start, end := 0, len(numbers)-1
	mid := (start+end)/2
	for start < end{
		mid = (start+end)/2
		if numbers[mid] < numbers[end] {
			//最小数在前半段
			end = mid
		}else if (numbers[mid]>numbers[end]){
			//最小数在后半段
			start = mid+1
		}else {
			end--
		}
	}
	return numbers[start]
}


//leetcode submit region end(Prohibit modification and deletion)
