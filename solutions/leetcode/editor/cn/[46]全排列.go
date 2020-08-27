package main
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。 
//
// 示例: 
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//] 
// Related Topics 回溯算法 
// 👍 857 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	//var ans [][]int
	ans := make([][]int,0,len(nums)*len(nums))
	t := make([]int,0,len(nums))
	ans = trace(nums,t,ans)
	return ans
}
func trace(nums []int, t []int,ans [][]int)[][]int{
	if len(nums) == len(t) {
		ct := make([]int,len(t),len(t))
		copy(ct,t)
		ans = append(ans, ct)
		return ans
	}
	for _,e := range nums{
		if isContain(t,e) {
			// t 已经有了 e ，跳过
			continue
		}
		t = append(t,e)
		ans = trace(nums,t,ans)
		t = remove(t,e)
	}
	return ans
}


func isContain(nums []int, n int) bool {
	for _,e :=range nums{
		if e == n{
			return true
		}
	}
	return false
}

func remove(t []int,n int) []int {
	for i,e := range t{
		if n == e{
			return append(t[:i],t[i+1:]...)
		}
	}
	return t
}

//leetcode submit region end(Prohibit modification and deletion)
