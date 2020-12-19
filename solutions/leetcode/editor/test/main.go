package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func findMaxForm(strs []string, m int, n int) int {
	/*
			0-1 背包问题 找出最大子集

			状态：m 和 n 和 str[i] 其实是有三种状态的，
				但由于使用了状态压缩，去掉了str[i]
			dp[m][n] : 满足条件m n的最大子集
			dp[0][0] = 0
			k= str[i]含0数 z=str[i]含1数
		    dp[m][n] = max{dp[m][n],dp[m-k][n-z]+1}

		难点：由于状态压缩，要对m和n，从大到小遍历
		在状态压缩时，为了避免后面的值被反复更新，要时候倒序遍历
	*/
	dp := make([][]int, 0)
	for i := 0; i <= m; i++ {
		dpi := make([]int, n+1)
		dp = append(dp, dpi)
	}

	for _, str := range strs {
		k, z := gen01(str)
		for i := m; i >= k; i-- {
			for j := n; j >= z; j-- {
				dp[i][j] = max(dp[i][j], dp[i-k][j-z]+1)
			}
		}
	}
	return dp[m][n]
}
func gen01(s string) (int, int) {
	k, z := 0, 0
	for _, e := range s {
		if string(e) == "0" {
			k++
		} else {
			z++
		}
	}
	return k, z
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
