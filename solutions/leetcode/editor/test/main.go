package main

import (
	"fmt"
)

func main() {
	fmt.Println(distinctSubseqII("lee"))
	fmt.Println(distinctSubseqII("abc"))
	/*
		a b       b
		1 1+2     3+(4-3)
	*/
	//fmt.Println(distinctSubseqII("aaa"))
}
func distinctSubseqII(S string) int {
	/*
		 // dp[i] 表示 S[0..i) 中有多少个**包含空序列**的不重复子序列
			1 + 2 + 4 + 8

			a b         a
			2 4     	6
			a a         a
			2 4-2   	4
	*/
	//int MOD = 1_000_000_007;
	N := len(S)
	dp := make([]int, N+1)
	dp[0] = 0

	last := make(map[uint8]int)

	for i := 0; i < N; i++ {
		x := S[i]
		t := last[x]
		dp[i+1] = dp[i]*2 + 1 - t
		last[x] = dp[i+1] - dp[i]
	}
	//dp[N]--
	//if (dp[N] < 0) dp[N] += MOD;
	return dp[N]
}
