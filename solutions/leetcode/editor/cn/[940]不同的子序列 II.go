package main
//给定一个字符串 S，计算 S 的不同非空子序列的个数。 
//
// 因为结果可能很大，所以返回答案模 10^9 + 7. 
//
// 
//
// 示例 1： 
//
// 输入："abc"
//输出：7
//解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
// 
//
// 示例 2： 
//
// 输入："aba"
//输出：6
//解释：6 个不同的子序列分别是 "a", "b", "ab", "ba", "aa" 以及 "aba"。
// abc
//a b c ab ac bc abc
// 示例 3： 
//
// 输入："aaa"
//输出：3
//解释：3 个不同的子序列分别是 "a", "aa" 以及 "aaa"。
// 
//
// 
//
// 
//
// 提示： 
//
// 
// S 只包含小写字母。 
// 1 <= S.length <= 2000 
// 
//
// 
//
// 
// Related Topics 动态规划 
// 👍 75 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func distinctSubseqII(S string) int {
	/*
		 dp[i+1] = 2*dp[i] +1
	dp[i]=2*dp[i-1]+1-重复计算的子序列数量
	而重复计算的子序列数量就是dpi-1
	至于为什么乘2，可以理解为如abc，dp[2]可以看作dp[1]的各种情况和dp[1]的各种情况后面都加了个c
	*/
	//int MOD = 1_000_000_007;
	N := len(S)
	dp := make([]int,N+1)
	dp[0] = 0;

	//last := make(map[uint8]int);

	for i:=0;i<N;i++{
		x := S[i]
		//t := last[x]
		dp[i+1] = dp[i]*2+1 - t
		for j:=i;j>=0;j++{
			if S[j] == x{
				if j > 0 {
					dp[i+1] = dp[i+1] - dp[j]  - 1
				}else {
					dp[i+1] = dp[i+1] - 1
				}
			}
		}
	}
	return dp[N];
}
//leetcode submit region end(Prohibit modification and deletion)
