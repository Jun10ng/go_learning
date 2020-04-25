package ex4_9

import (
	"fmt"
)

/*
 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。
在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
*/

func wordFreq(words []string)  {
	count := make(map[string]int)
	for _,word :=range words{
		count[word] +=1
	}
	for k,v := range count {
		fmt.Printf("%v \t %v \n",k,v)
	}
}