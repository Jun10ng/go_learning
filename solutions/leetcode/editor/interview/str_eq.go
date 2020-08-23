package interview

//递归
func isEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	len := len(s1)
	if (len % 2) == 1 {
		// 字符串长度为奇数的情况下，只有完全相同才能相等
		return false
	}
	mid := len / 2
	if isEqual(s1[:mid], s2[mid:]) && isEqual(s1[mid:], s2[:mid]) {
		return true
	}
	if isEqual(s1[mid:], s2[mid:]) && isEqual(s1[:mid], s2[:mid]) {
		return true
	}
	return false
}
