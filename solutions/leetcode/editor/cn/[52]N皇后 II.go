package main
//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 
//
// 上图为 8 皇后问题的一种解法。 
//
// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。 
//
// 示例: 
//
// 输入: 4
//输出: 2
//解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
// 
//
// 
//
// 提示： 
//
// 
// 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或 N
//-1 步，可进可退。（引用自 百度百科 - 皇后 ） 
// 
// Related Topics 回溯算法 
// 👍 142 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	var ans int
	cs := &[]string{}

	var line string
	for i := 0; i < n; i++ {
		line += "."
	}
	for i := 0; i < n; i++ {
		*cs = append(*cs, line)
	}

	//
	trace(&ans,cs, 0, 0, n)
	return ans
}
func trace(ans *int, cs *[]string, row, col, n int) {
	if row == n {
		*ans++
		return
	}

	for i := 0; i < n; i++ {
		if !Ok(cs, row, i) {
			continue
		}
		(*cs)[row] = (*cs)[row][:i] + "Q" + (*cs)[row][i+1:]
		trace(ans, cs, row+1, 0, n)
		(*cs)[row] = (*cs)[row][:i] + "." + (*cs)[row][i+1:]
	}

}

func Ok(cs *[]string, row, col int) bool {
	n := len(*cs)
	for i := 0; i < n; i++ {
		if (*cs)[i][col] == 'Q' {
			return false
		}
	}
	//左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if (*cs)[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	//右上
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if (*cs)[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
