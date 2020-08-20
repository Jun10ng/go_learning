package main
//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。 
//
// 
//
// 示例 1： 
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
// 
//
// 示例 2： 
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
// 
//
// 
//
// 限制： 
//
// 
// 0 <= matrix.length <= 100 
// 0 <= matrix[i].length <= 100 
// 
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/ 
// Related Topics 数组 
// 👍 126 👎 0


var m,n int
var si,sj,ei,ej int //屏障
func spiralOrder(matrix [][]int) []int {
	m = len(matrix)
	if m==0{
		return nil
	}
	n = len(matrix[0])
	if n==0{
		return nil
	}
	si,sj,ei,ej = 0,0,m-1,n-1
	c,i,j := 0,0,0
	ret := make([]int,0)

	for c < m*n{
		// 从左到右
		for j=sj;j<=ej&&c<m*n;j++{
			ret = append(ret, matrix[si][j])
			c++
		}
		for i=si+1;i<=ei&&c<m*n;i++{
			ret = append(ret, matrix[i][ej])
			c++
		}
		for j=ej-1;j>=sj&&c<m*n;j--{
			ret = append(ret, matrix[ei][j])
			c++
		}
		for i=ei-1;i>=si+1&&c<m*n;i--{
			ret = append(ret,matrix[i][sj])
			c++
		}
		si+=1
		sj+=1
		ej-=1
		ei-=1
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
