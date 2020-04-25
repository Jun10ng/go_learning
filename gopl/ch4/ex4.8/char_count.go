package ex4_8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

/*
**练习 4.8：** 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。

令人惊讶的是 汉字也算是letter
*/

func charCount(s string)  {
	var letterCount,numberCount,otherCount int
	// 输入
	//将 s 作为 输入
	in := bufio.NewReader(strings.NewReader(s))
	for  {
		// 以 rune 为单位读取
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		// 无效码点
		if r == unicode.ReplacementChar && n == 1 {
			//invalid++
			continue
		}
		if unicode.IsNumber(r) {
			numberCount++
			fmt.Printf("%v is number\n",r)
		}else if unicode.IsLetter(r){
			letterCount++
			fmt.Printf("%v is letter\n",r)
		}else {
			otherCount++
		}

	}
	fmt.Printf("letter: %v,\t number,%v\t other: %v",letterCount,numberCount,otherCount)
}