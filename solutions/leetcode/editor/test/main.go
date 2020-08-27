package main

import (
	"fmt"
	"sort"
)

//给定二叉树: [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
//

//* Definition for a binary tree node.
func main() {
	nums := []int{1, 1, 2}
	ans := permuteUnique(nums)
	fmt.Print(ans)
}

func permuteUnique(nums []int) [][]int {
	//var ans [][]int
	sort.Ints(nums)
	ans := make([][]int, 0, len(nums)*len(nums))
	t := make([]int, 0, len(nums))
	used := make([]bool, len(nums), len(nums))
	trace(nums, t, used, &ans)
	return ans
}
func trace(nums []int, t []int, used []bool, ans *[][]int) {
	if len(nums) == len(t) {
		ct := make([]int, len(t), len(t))
		copy(ct, t)
		*ans = append(*ans, ct)
		return
	}
	for i, e := range nums {
		if used[i] {
			continue
		}
		if i != 0 && used[i-1] && e == nums[i-1] {
			continue
		}
		used[i] = true
		t = append(t, e)
		trace(nums, t, used, ans)
		used[i] = false
		t = removeLast(t)
	}
	return
}

func isContains(t []int, n int) bool {
	for _, e := range t {
		if e == n {
			return true
		}
	}
	return false
}

func removeLast(t []int) []int {
	return t[:len(t)-1]
}
