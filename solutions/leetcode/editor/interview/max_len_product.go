package interview

// 解析：abcbbc 状态压缩为 111 （前三个字母都存在），只要两个字符串状态压缩后，&操作后等于0，说明没有相同字母存在
func maxLenPdt(s string) int {

	maxChar := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] > maxChar {
			maxChar = s[i]
		}
	}
	// s 是由 字母表前 n 位字母组成的。
	n := int(maxChar - uint8('a') + 1)
	z := 1 << n
	//dp[x] 表示 状态为x的字符串的最长长度
	dp := make([]int, z)

	//初始化dp数组，记录每种状态字符串的最长长度
	for i := 0; i < len(s); i++ {
		x := 0
		for j := i; j < len(s); j++ {
			f := 1 << (s[j] - 'a')
			x |= f
			dp[x] = max(dp[x], j-i+1)
		}
	}

	//此时dp[i] 表示集合为s的最长长度
	for i := 1; i < z; i++ {
		for j := 0; j < n; j++ {
			y := 1 << j
			if (i & y) != 0 {
				dp[i] = max(dp[i], dp[i-y])
			}
		}
	}
	//遍历所有状态与其相补的状态的字符串，记录最大乘积
	pdt := 0
	for i := 0; i < z; i++ {
		pdt = max(pdt, dp[i]*dp[z-i-1])
	}
	return pdt
}
