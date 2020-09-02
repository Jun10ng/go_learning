package main

//我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//
// 
//
// 示例: 
//
// 输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。 
//
// 说明: 
//
// 
// 1 是丑数。 
// n 不超过1690。 
// 
//
// 注意：本题与主站 264 题相同：https://leetcode-cn.com/problems/ugly-number-ii/ 
// Related Topics 数学 
// 👍 61 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int{
	if n <= 1{
		return n
	}
	q := make([]int,0,n+1)
	q = append(q,1)

	nxt2 := 0
	nxt3 := 0
	nxt5 := 0

	tmp := 0
	for len(q)<n{
		tmp = min3(2*q[nxt2],3*q[nxt3],5*q[nxt5])
		q = append(q,tmp)

		if tmp == 2*q[nxt2] {nxt2++}
		if tmp == 3*q[nxt3] {nxt3++}
		if tmp == 5*q[nxt5] {nxt5++}
	}
	return q[len(q)-1]
}

func min2(a,b int) int {
	if a < b{
		return a
	}
	return b
}

func min3(a,b,c int) int {
	tmp := min2(a,b)
	return min2(tmp,c)
}

//leetcode submit region end(Prohibit modification and deletion)
