package main
//给定一个字符串数组 arr，字符串 s 是将 arr 某一子序列字符串连接所得的字符串，如果 s 中的每一个字符都只出现过一次，那么它就是一个可行解。 
//
// 请返回所有可行解 s 中最长长度。 
//
// 
//
// 示例 1： 
//
// 输入：arr = ["un","iq","ue"]
//输出：4
//解释：所有可能的串联组合是 "","un","iq","ue","uniq" 和 "ique"，最大长度为 4。
// 
//
// 示例 2： 
//
// 输入：arr = ["cha","r","act","ers"]
//输出：6
//解释：可能的解答有 "chaers" 和 "acters"。
// 
//
// 示例 3： 
//
// 输入：arr = ["abcdefghijklmnopqrstuvwxyz"]
// 输出：26
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr.length <= 16 
// 1 <= arr[i].length <= 26 
// arr[i] 中只含有小写英文字母 
// 
// Related Topics 位运算 回溯算法 
// 👍 61 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxLength(arr []string) int {
	if len(arr)==0 {
		return 26
	}
	ans := 0

	trace := func(cur string,begin int) {}
	trace = func(cur string,begin int) {
		if begin > len(arr) {
			return
		}
		if len(cur) > ans {
			ans = len(cur)
		}

		for i:=begin;i<len(arr);i++ {
			if !isOk(arr[i],cur) {
				continue
			}
			tmpCnt := len(arr[i])
			cur = cur+arr[i]
			trace(cur,i+1)
			cur = cur[:len(cur)-tmpCnt]
		}
	}
	trace("",0)
	return ans
}

func isOk(arri,cur string) bool {
	arrmap := make(map[int32]bool,len(arri))
	for _,ai := range arri{
		if _,exist:=arrmap[ai];exist {
			return false
		}
		arrmap[ai]=true
	}
	for _,ci :=range cur{
		if _,exist:=arrmap[ci];exist {
			return false
		}
	}
	return true

}
//leetcode submit region end(Prohibit modification and deletion)
