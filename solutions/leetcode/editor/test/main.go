package main

import (
	"fmt"
)

func main() {

	//fmt.Println(generateParenthesis(3))

	fmt.Println(minDistance("horse", "ros"))
}

type Mem [][]int

func (m Mem) Repalce(i, j int) int {
	return m[i-1][j-1]
}
func (m Mem) Insert(i, j int) int {
	return m[i][j-1] // ro  ros
}
func (m Mem) Delete(i, j int) int {
	return m[i-1][j] // ros ro
}
func minDistance(word1 string, word2 string) int {
	/*
		note！如果有两个string 入参，dp的状态定义一般都是二维，前ij各表示一个string
		dp[i][j] : word1前i个字符 和 word2前j个字符的最小操作次数

	*/
	//dp := [][]int{}
	dp := Mem{}

	for i := 0; i <= len(word1); i++ {
		dpi := make([]int, len(word2)+1)
		dp = append(dp, dpi)
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	//
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min3(dp.Delete(i, j), dp.Insert(i, j), dp.Repalce(i, j))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
func minimumTotal(triangle [][]int) int {
	cur := []int{}
	pre := []int{}
	for i, e := range triangle {
		pi := i - 1
		if pi < 0 {
			cur = append(cur, e[0])
		} else {
			for ei, ee := range e {
				if ei-1 < 0 {
					cur = append(cur, pre[ei]+ee)
				} else if ei == len(e)-1 {
					cur = append(cur, pre[ei-1]+ee)
				} else {
					cur = append(cur, min2(pre[ei], pre[ei-1])+ee)
				}
			}
		}
		pre = cur
		cur = make([]int, 0)
	}
	min := 9999999999
	for _, ei := range pre {
		min = min2(min, ei)
	}
	return min
}

func max2(a, b int) int {
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
	t := min2(a, b)
	return min2(t, c)
}
