package main
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。 
//
// 示例: 
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
// Related Topics 回溯算法 
// 👍 338 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	ans := &[][]int{}
	e := &[]int{}
	t := make([]int,0,n)
	for i:=1;i<=n;i++ {
		t = append(t,i)
	}
	trace(ans,e,&t,k,0)
	return *ans
}

func trace(ans *[][]int, e,t *[]int,k,begin int)  {
	if len(*e) == k {
		cp := make([]int,k)
		copy(cp,*e)
		*ans = append(*ans,cp)
		return
	}
	for i:= begin;i<len(*t);i++ {
		*e = append(*e,(*t)[i])
		trace(ans,e,t,k,i+1)
		*e = (*e)[:len(*e)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
