package ex4_3

/*
练习 4.3： 重写reverse函数，使用数组指针代替slice。
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
*/
func reverse(a *[5]int){
	for i,j := 0,len(*a)-1; i<j; i,j = i+1,j-1  {
		a[i], a[j] = a[j], a[i]
	}
}