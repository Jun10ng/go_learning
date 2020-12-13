package utils

import (
	"fmt"
	"testing"
)

var a, b int

func f() {
	a = 1 // w之前的写操作
	b = 2 // 写操作w
}

func g() {
	print(b) // 读操作r
	print(a) // ???
}

func TestMeme(t *testing.T) {
	go f() //g1
	//time.Sleep(time.Second)
	g() //g2
}

var s string
var done bool

func setup() {
	s = "hello, world"
	done = true
}

func TestMeme2(t *testing.T) {
	go setup()
	for !done {
	}
	print(s)
}

// slice 插入元素
func TestSliceInsert2(t *testing.T) {
	fmt.Println(Encode("12324@qq.com", true))
}
func SliceInsert(s []string, index int, value string) []string {
	rear := append([]string{}, s[index:]...)
	return append(append(s[:index], value), rear...)
}

func Encode(s string, isEmail bool) string {
	if isEmail {
		j := 0
		for ; j < len(s); j++ {
			if string(s[j]) == "@" {
				break
			}
		}

		if j < 2 {
			return s
		}
		x := ""
		for i := 0; i < j-2; i++ {
			x = x + "*"
		}

		return s[:2] + x + s[j:]

	}

	if len(s) <= 5 {
		return s
	}
	x := ""
	for i := 0; i < len(s)-5; i++ {
		x = x + "*"
	}
	return s[:3] + x + s[len(s)-2:]
}
