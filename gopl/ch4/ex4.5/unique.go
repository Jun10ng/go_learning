package ex4_5

/*
练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
*/

func unique(s []string) []string {
	k := 0
	for _,w :=range s{
		if s[k]==w {
			continue
		}
		k++
		s[k]=w
	}
	return s[:k+1]
}