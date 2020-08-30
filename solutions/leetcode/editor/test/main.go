package main

import (
	"fmt"
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
	fmt.Print(solveNQueens(4))
}
func solveNQueens(n int) [][]string {
	ans := &[][]string{}
	cs := &[]string{}

	var line string
	for i := 0; i < n; i++ {
		line += "."
	}
	for i := 0; i < n; i++ {
		*cs = append(*cs, line)
	}

	//
	trace(ans, cs, 0, 0, n)
	return *ans
}

func trace(ans *[][]string, cs *[]string, row, col, n int) {
	if row == n {
		cp := make([]string, n)
		copy(cp, *cs)
		*ans = append(*ans, cp)
		return
	}

	for i := 0; i < n; i++ {
		if !Ok(cs, row, i) {
			continue
		}
		(*cs)[row] = (*cs)[row][:i] + "Q" + (*cs)[row][i+1:]
		trace(ans, cs, row+1, 0, n)
		(*cs)[row] = (*cs)[row][:i] + "." + (*cs)[row][i+1:]
	}

}

func Ok(cs *[]string, row, col int) bool {
	n := len(*cs)
	for i := 0; i < n; i++ {
		if (*cs)[i][col] == 'Q' {
			return false
		}
	}
	//左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if (*cs)[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	//右上
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if (*cs)[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}
