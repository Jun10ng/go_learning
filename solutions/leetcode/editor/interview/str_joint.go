package interview

// 利用KMP 求最小循环节
func strJoint(s string) int {
	next := make([]int, len(s)+1)
	next[0] = -1
	len := len(s)
	i, j := 0, -1
	for i < len {
		for j != -1 && s[i] != s[j] {
			j = next[j]
		}
		i += 1
		j += 1
		next[i] = j
	}

	var ans int
	ans = len - next[len]
	if ans != len && len%ans == 0 {
		return 0
	} else {
		return ans - (len - len/ans*ans)
	}
	return 0
}
