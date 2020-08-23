package interview

import (
	"fmt"
	"math"
	"strconv"
)

// 全排列 + 数学期望
func expect(n int, p float64) float64 {
	var p1, p2, ans float64
	c := math.Log(1.0)

	for i := 0; i < n+1; i++ {
		//两种情况，要么那天做的是第一套，要么是第二套
		p1 += math.Log(p)
		p2 += math.Log(1 - p)
	}
	// i 是天数，没题做的天数最小是第n+1天，最大是第2n+1天
	for i := n + 1; i <= 2*n+1; i++ {
		ans = ans + (math.Exp(c+p1)+math.Exp(c+p2))*float64(2*n-i+1)
		c += math.Log(float64(i) / float64(i-n))
		p1 += math.Log(1 - p)
		p2 += math.Log(p)
	}
	ans, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", ans), 64)
	return ans
}
