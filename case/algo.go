package main

import "fmt"

/*
多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。


示例 1：

输入：[3,2,3]
输出：3
示例 2：

输入：[2,2,1,1,1,2,2]
输出：2
*/

func find(nums []int) (int, bool) {
	ptr := 0
	cnt := 1
	maxCnt := 0

	for i := 1; i < len(nums); i++ {
		e := nums[i]
		if e == nums[ptr] {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ptr = i + 1
			cnt = 1
			i++
		}
		maxCnt = max2(cnt, maxCnt)
	}
	if cnt == maxCnt {
		return 0, false
	}
	return nums[ptr], true
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(find([]int{2, 1, 2, 1, 3, 3}))
}
