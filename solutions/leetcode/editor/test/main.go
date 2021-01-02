package main

import (
	"fmt"
)

func main() {
	m := []int{4, 6, 7, 7} // 1357 1347
	fmt.Println(findSubsequences(m))
}

func findSubsequences(nums []int) [][]int {
	ans := [][]int{}
	doFind := func(path []int, i int) {}
	doFind = func(path []int, i int) {
		if len(path) > 1 {
			cpy := make([]int, len(path))
			copy(cpy, path)
			ans = append(ans, cpy)
		}
		vis := make(map[int]bool)
		for j := i; j < len(nums); j++ {
			if vis[nums[j]] {
				continue
			}
			if len(path) == 0 || nums[j] >= path[len(path)-1] {
				vis[nums[j]] = true
				path = append(path, nums[j])
				doFind(path, j+1)
				path = path[:len(path)-1]
			}
		}
	}

	doFind([]int{}, 0)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	tmp := min2(a, b)
	return min2(tmp, c)
}
