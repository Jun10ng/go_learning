package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

func verifyPostorder(postorder []int) bool {
	return verify(postorder, 0, len(postorder)-1)
}

func verify(p []int, s int, e int) bool {
	if s >= e {
		return true
	}
	root := p[e]
	i := s
	for p[i] < root {
		i++
	}
	tmp := i
	//从i开始的，有比root小的值，是不合理的。
	for tmp < e {
		if p[tmp] < root {
			return false
		}
		tmp++
	}
	return verify(p, s, i-1) && verify(p, i, e-1)
}
func main() {
	s := []int{1, 3, 2, 6, 5}
	fmt.Print(verifyPostorder(s))
}
