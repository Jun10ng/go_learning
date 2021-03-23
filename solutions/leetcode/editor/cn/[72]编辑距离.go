package main
//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。 
//
// 你可以对一个单词进行如下三种操作： 
//
// 
// 插入一个字符 
// 删除一个字符 
// 替换一个字符 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
// 
//
// 示例 2： 
//
// 
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
// 
//
// 
//
// 提示： 
//
// 
// 0 <= word1.length, word2.length <= 500 
// word1 和 word2 由小写英文字母组成 
// 
// Related Topics 字符串 动态规划 
// 👍 1484 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type Mem [][]int

func (m Mem) Repalce(i,j int) int{
	return m[i-1][j-1]
}
func (m Mem) Insert(i,j int) int{
	return m[i][j-1] // ro  ros
}
func (m Mem) Delete(i,j int) int{
	return m[i-1][j] // ros ro
}
func minDistance(word1 string, word2 string) int {
	/*
		note！如果有两个string 入参，dp的状态定义一般都是二维，前ij各表示一个string
		dp[i][j] : word1前i个字符 和 word2前j个字符的最小操作次数

	*/
	//dp := [][]int{}
	dp := Mem{}

	for i:=0;i<=len(word1);i++{
		dpi := make([]int,len(word2)+1)
		dp = append(dp,dpi)
		dp[i][0] = i
	}
	for j:=0;j<=len(word2);j++ {
		dp[0][j] = j
	}

	//
	for i:=1;i<=len(word1);i++{
		for j:=1;j<=len(word2);j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else {
				dp[i][j] = 1+ min3(dp.Delete(i,j),dp.Insert(i,j),dp.Repalce(i,j))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}


func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a, b,c int) int {
	t := min2(a,b)
	return min2(t,c)
}
//leetcode submit region end(Prohibit modification and deletion)
