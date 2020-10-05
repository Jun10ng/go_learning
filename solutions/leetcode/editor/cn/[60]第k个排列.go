package main

import "strconv"

//给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
//
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下： 
//
// 
// "123" 
// "132" 
// "213" 
// "231" 
// "312" 
// "321" 
// 
//
// 给定 n 和 k，返回第 k 个排列。 
//
// 说明： 
//
// 
// 给定 n 的范围是 [1, 9]。 
// 给定 k 的范围是[1, n!]。 
// 
//
// 示例 1: 
//
// 输入: n = 3, k = 3
//输出: "213"
// 
//
// 示例 2: 
//
// 输入: n = 4, k = 9
//输出: "2314"
// 
// Related Topics 数学 回溯算法 
// 👍 312 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

//康拓展开法

var(
	jc = []int{1,2,6,24,120,720,5040,40320,362880,3628800}
)
func getPermutation(n int, k int) string {
	// 输入: n = 3, k = 3
	//输出: "213"
	ans := ""
	//arr := make([]int,n)
	rest := make([]int,n)// 1 2 3

	for i := 0; i <n ; i++ {
		rest[i] = i+1
	}
	k = k-1

	for i := 1; i <n ; i++ {
		m := k/jc[n-i-1]
		k = k%jc[n-i-1]
		ans = ans + strconv.Itoa(rest[m])
		rest = append(rest[:m],rest[m+1:]...)
	}
	ans= ans+strconv.Itoa(rest[0])
	return ans
}



//回溯法
//func getPermutation(n int, k int) string {
//	ans := &[]string{}
//	t := ""
//	e := ""
//	for i:=1;i<=n;i++ {
//		//"1234"
//		t = t + strconv.Itoa(i)
//	}
//	used := make([]bool,len(t))
//	trace(ans,used,t,e,k)
//	return (*ans)[len(*ans)-1]
//}
//func trace(ans *[]string,used []bool, t,e string, k int)  {
//	if len(e) == len(t){
//		*ans = append((*ans),e)
//		return
//	}
//
//	for i:=0;i<len(t)&&len(*ans)<k;i++ {
//		if used[i] {
//			continue
//		}
//		used[i]=true
//		e = e + string(t[i])
//		trace(ans,used,t,e,k)
//		used[i] = false
//		e = e[:len(e)-1]
//	}
//}
//leetcode submit region end(Prohibit modification and deletion)
