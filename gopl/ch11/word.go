package ch11

import "unicode"

// 判断一个字符串是否前后相等
func IsPalindrome(s string) bool {
	var r []rune
	// 去除空格 并 转化为rune字符

	for _,e := range s{
		if unicode.IsLetter(e){
			r = append(r,unicode.ToLower(e))
		}
	}
	for i := range r {
		if r[i] !=r[len(r)-1-i] {
			return false
		}
	}
	return true
}