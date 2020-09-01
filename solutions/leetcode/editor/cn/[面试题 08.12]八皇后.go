package main
//设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整
//个棋盘的那两条对角线。 
//
// 注意：本题相对原题做了扩展 
//
// 示例: 
//
//  输入：4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释: 4 皇后问题存在如下两个不同的解法。
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
// Related Topics 回溯算法 
// 👍 39 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	ans := &[][]string{}
	cs := &[]string{}
	var line string
	for i := 0; i < n; i++ {
		line += "."
	}
	for i := 0; i < n; i++ {
		*cs = append(*cs, line)
	}

	trace(ans, cs,0,n)
	return *ans
}
func trace(ans *[][]string, cs *[]string,row,n int )  {
	if row==n {
		cp := make([]string,n)
		copy(cp,*cs)
		*ans = append(*ans,cp)
		return
	}
	for i:= 0;i<n;i++ {
		if Ok(cs,row,i) {
			(*cs)[row] = (*cs)[row][:i] + "Q" + (*cs)[row][i+1:]
			trace(ans, cs, row+1, n)
			(*cs)[row] = (*cs)[row][:i] + "." + (*cs)[row][i+1:]
		}else {
			continue
		}
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
