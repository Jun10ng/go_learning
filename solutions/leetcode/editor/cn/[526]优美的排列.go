package main
//假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N) 满足如下两个条件中的一个，
//我们就称这个数组为一个优美的排列。条件： 
//
// 
// 第 i 位的数字能被 i 整除 
// i 能被第 i 位上的数字整除 
// 
//
// 现在给定一个整数 N，请问可以构造多少个优美的排列？ 
//
// 示例1: 
//
// 
//输入: 2
//输出: 2
//解释: 
//
//第 1 个优美的排列是 [1, 2]:
//  第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
//  第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除
//
//第 2 个优美的排列是 [2, 1]:
//  第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
//  第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
// 
//
// 说明: 
//
// 
// N 是一个正整数，并且不会超过15。 
// 
// Related Topics 回溯算法 
// 👍 98 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

/*
这题不难 但是使用golang来做就有点恶心，因为golang的slice的特性，很难保持隔离性，
但是可以使用visited map来判断一个数是否有使用过。这样就避免了使用slice
但是，我就是不想做了，太恶心。

*/
func countArrangement(N int) int {
	cnt := 0
	rest := make([]int,0)
	for i:=1;i<=N;i++ {
		rest = append(rest,i)
	}

	/*
		// 第 i 位的数字能被 i 整除
		// i 能被第 i 位上的数字整除
			2 [1,2] [2,1]
			3 [1,2,3] [2,1,3] [3,2,1]
	*/
	trace := func(rest []int,idx int) {}
	trace = func(rest []int,idx int) {
		if len(rest)==0{
			cnt++
		}
		for i:=0;i<len(rest);i++{
			if isOk(idx,rest[i]) {
				tmp := rest[i]
				rest = append(rest[:i],rest[i+1:]...)
				trace(rest,idx+1)
				tmpSlice :=append(rest[:i],tmp)
				rest = append(tmpSlice,rest[i:]...)
			}
		}
	}

	trace(rest,1)
	return cnt
}

func isOk(idx,e int)bool  {
	if idx%e == 0 || e%idx ==0{
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
