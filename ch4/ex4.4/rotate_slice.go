package ex4_4
/*
练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
*/

//返回 将target切片 前 n 个元素选择的新切片
func rotate(t []int,n int)[]int  {
	/*
		这个地方我起初写成
		tmp := t[:n]
		测试是发现始终输出：[3,4,5,3,4]
		才想起来这是 浅拷贝，是浮动的
		正确使用要使用深拷贝去给tmp赋值
	*/
	tmp := make([]int,n)
	copy(tmp,t[:n])
	copy(t[:len(t)-n],t[n:])
	copy(t[len(t)-n:],tmp[:])
	return t
}