package main

import "fmt"

//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果
//一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。 
//
// [["a","b","c","e"], 
//["s","f","c","s"], 
//["a","d","e","e"]] 
//
// 但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。 
//
// 
//
// 示例 1： 
//
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "A
//BCCED"
//输出：true
// 
//
// 示例 2： 
//
// 输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
// 
//
// 提示： 
//
// 
// 1 <= board.length <= 200 
// 1 <= board[i].length <= 200 
//


//leetcode submit region begin(Prohibit modification and deletion)
var v = [][]int{{-1,0},{1,0},{0,-1},{0,1}}

func exist(board [][]byte, word string) bool {
	idx := 0
	for i:=0;i<len(board);i++ {
		for j := 0;j<len(board[0]);j++ {
			if trace(board,word,idx,i,j) {
				return true
			}
		}
	}
	return false
}

func trace(board [][]byte, word string, idx,i,j int) bool {
	if idx == len(word) {
		return true
	}

	if(i<0||j<0||i>=len(board)||j>=len(board[0])){
		return false
	}
	if board[i][j]==word[idx]{
		tmp := board[i][j]
		board[i][j] = '.'
		for k:=0;k<4;k++ {
			i,j = i+v[k][0],j+v[k][1]
			if trace(board,word,idx+1,i,j){
				return true
			}
			i,j = i-v[k][0],j-v[k][1]
		}
		board[i][j] = tmp
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
