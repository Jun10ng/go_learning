package main
//如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的： 
//
// 
// n >= 3 
// 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2} 
// 
//
// 给定一个严格递增的正整数数组形成序列，找到 A 中最长的斐波那契式的子序列的长度。如果一个不存在，返回 0 。 
//
// （回想一下，子序列是从原序列 A 中派生出来的，它从 A 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3
//, 4, 5, 6, 7, 8] 的一个子序列） 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入: [1,2,3,4,5,6,7,8]
//输出: 5
//解释:
//最长的斐波那契式子序列为：[1,2,3,5,8] 。
// 
//
// 示例 2： 
//
// 输入: [1,3,7,11,12,14,18]
//输出: 3
//解释:
//最长的斐波那契式子序列有：
//[1,11,12]，[3,11,14] 以及 [7,11,18] 。
// 
//
// 
//
// 提示： 
//
// 
// 3 <= A.length <= 1000 
// 1 <= A[0] < A[1] < ... < A[A.length - 1] <= 10^9 
// （对于以 Java，C，C++，以及 C# 的提交，时间限制被减少了 50%） 
// 
// Related Topics 数组 动态规划 
// 👍 156 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func lenLongestFibSubseq(arr []int) int {
	/*
		状态
		dp[i][j] 以 arr[i] arr[j] 结尾的最长
		dp[i][j] = max(dp[i][j],dp[k][i]+1) arr[k]+arr[i] = arr[j] k<i
	*/
	dp := [][]int{}
	for i:=0;i<=len(arr);i++{
		dpi := make([]int,len(arr)+1)
		for ei,_:=range dpi{
			dpi[ei] = 2
		}
		dp = append(dp,dpi)
	}
	ret := 0
	/*
		for i:=0;i<len(arr)-1;i++{
			for j:=i+1;j<len(arr);j++ {
				for k:=0;k<i;k++{
					if arr[i]+arr[k] == arr[j] {
						dp[i+1][j+1] = max2(dp[i+1][j+1],dp[k+1][i+1]+1)
						ret = max2(ret,dp[i+1][j+1])
					}
				}
			}
		}
		优化为下方代码
	*/
	pathMap := make(map[int]int,len(arr)) // pathMap为arr的逆
	for i,e := range arr{
		pathMap[e] = i+1
	}
	for i:=0;i<len(arr)-1;i++{
		for j:=i+1;j<len(arr);j++ {
			if k,ok :=pathMap[arr[j]-arr[i]];ok&&k<i+1{
				dp[i+1][j+1] = max2(dp[i+1][j+1],dp[k][i+1]+1)
				ret = max2(ret,dp[i+1][j+1])
			}
		}
	}

	return ret

}
// 回溯法
//func lenLongestFibSubseq(arr []int) int {
//	/*
//		回溯法
//		[1,2,3,5,8]
//	*/
//	ret := 0
//	cur := 2
//	trace := func(int,int,int) {}
//	trace = func(idx,a,b int) {
//		for i:= idx+1;i<len(arr);i++ {
//			tmp := cur
//			if arr[i] == a+b {
//				cur++
//				ret = max2(cur,ret)
//				trace(i,b,a+b)
//			}
//			cur = tmp
//		}
//	}
//
//	for i:= 0;i<len(arr)-2;i++{
//		for j:=i+1;j<len(arr);j++{
//			trace(j,arr[i],arr[j])
//		}
//		//cur = 2
//	}
//	return ret
//}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
