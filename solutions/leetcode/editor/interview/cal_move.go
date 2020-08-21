package interview

//一个乱序数组，求要移动几次数字才能使得其有序
//解析：其实是最长上升子序列的反向问题，结果 = 数组长度 - 最长上升子序列长度。
func calMove(q []int) int {
	if len(q) == 0 {
		return 0
	}
	dp := make([]int, 0)

	for i := 0; i < len(q); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if q[j] < q[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return len(q) - maxOf(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxOf(q []int) int {
	ret := q[0]
	for i := 1; i < len(q); i++ {
		if q[i] > ret {
			ret = q[i]
		}
	}
	return ret
}
