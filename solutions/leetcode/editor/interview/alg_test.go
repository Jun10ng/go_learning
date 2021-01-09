package interview

import (
	"fmt"
	"testing"
)

//计算移动次数
func TestCalMove(t *testing.T) {
	q := []int{1, 101, 99, 88, 77, 66, 55, 44, 33, 22, 2}
	fmt.Print(calMove(q))
}

// 判断字符串是否相等
func TestStrEq(t *testing.T) {
	fmt.Print(isEqual("aaaa", "aaaaaa"))
	//fmt.Print(isEqual("aabb","abab"))
}

//计算字符串相乘最大长度
func TestMaxPdt(t *testing.T) {
	s := "fccecccaafaadaabgbb"
	fmt.Print(maxLenPdt(s))
}

//计算数学题的期望
func TestExpct(t *testing.T) {
	fmt.Print(expect(500, 0.56))
}

//合并瓜子
func TestGuazi(t *testing.T) {
	g := []int{}
	fmt.Print(guazi(g))
}

//求t字符串长度
func TestStrJoint(t *testing.T) {
	s := "abcdabcdabc"
	fmt.Print(strJoint(s))
}

func TestMoveNums(t *testing.T) {
	n := []int{-1, 2, 0, -3, 0, 4, -5}
	fmt.Println(moveNum(n))
}
