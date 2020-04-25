package ex4_9

import (
	"bufio"
	"os"
	"testing"
)

func TestWordFreq(t *testing.T)  {
	cin := bufio.NewScanner(os.Stdin)
	cin.Split(bufio.ScanWords)
	s:=make([]string,10)
	for cin.Scan(){
		s=append(s,cin.Text())
	}

	wordFreq(s)
}