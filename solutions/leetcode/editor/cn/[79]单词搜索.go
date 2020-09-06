package main
//给定一个二维网格和一个单词，找出该单词是否存在于网格中。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。 
//
// 
//
// 示例: 
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false 
//
// 
//
// 提示： 
//
// 
// board 和 word 中只包含大写和小写英文字母。 
// 1 <= board.length <= 200 
// 1 <= board[i].length <= 200 
// 1 <= word.length <= 10^3 
// 
// Related Topics 数组 回溯算法 
// 👍 544 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var v = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
func exist(board [][]byte, word string) bool {
	var ans bool
	e := ""+string(word[0])
	for i:=0;i<len(board);i++ {
		for j:=0;j<len(board[0]);j++ {
			if board[i][j]==word[0] {
				tmp := board[i][j]
				board[i][j] = '.'
				trace(&ans,&board,e,word,i,j)
				if ans {
					return ans
				}
				board[i][j] = tmp
			}
		}
	}
	return false
}
func trace(ans *bool,board *[][]byte, e,word string,i,j int )  {
	if e == word {
		*ans = true
		return
	}
	if len(e)>len(word){
		return
	}

	//
	for _,vi := range v{
		i, j = i+vi[0],j+vi[1]
		if !(*ans) &&i>=0 && j>=0 && i < len(*board) && j < len((*board)[0]) && (*board)[i][j] == word[len(e)]{
			e = e + string((*board)[i][j])
			(*board)[i][j] =  '.'
			trace(ans,board,e,word,i,j)
			(*board)[i][j] =  e[len(e)-1]
			e = e[:len(e)-1]
		}
		i, j = i-vi[0],j-vi[1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
