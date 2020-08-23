package interview

const (
	N   = 100
	Max = 10000 * N
)

// 操场石子合并改编题
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var dp [N][N]int

func guazi(g []int) int {
	n := len(g)
	s := make([]int, 0)
	s = append(s, 0)
	for i := 1; i <= n; i++ {
		s = append(s, g[i-1])
		s[i] += s[i-1]
	}

	// 这是按照组合堆数的方式进行遍历，一开始组合相邻的2堆、3堆、4堆...
	for len := 2; len <= n; len++ {
		for i := 1; i <= n-len+1; i++ {
			j := i + len - 1
			dp[i][j] = Max
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+s[j]-s[i-1])
			}
		}
	}

	return dp[1][n]
}
