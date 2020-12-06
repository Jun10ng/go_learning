package main

import "strconv"

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 示例 1: 
//
// 输入: 121
//输出: true
// 
//
// 示例 2: 
//
// 输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 
//
// 示例 3: 
//
// 输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。
// 
//
// 进阶: 
//
// 你能不将整数转为字符串来解决这个问题吗？ 
// Related Topics 数学 
// 👍 1329 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	strX := strconv.Itoa(x)

	// normal
	/*
		12321
	*/
	prt1 := 0     // 从前往后的指针
	prt2 := len(strX)-1 //

	for ;prt1<=prt2;{
		if strX[prt1]!=strX[prt2] {
			return false
		}
		prt1++
		prt2--
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)

















