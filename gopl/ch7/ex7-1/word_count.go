package ch7

import (
	"bufio"
)

//计算单词和行数
type WordCounter struct {
	Word int
	Line int
}

func (wc *WordCounter) write(p []byte) (int, error) {
	if p != nil {
		wc.Line = 1
	}
	advance, _, err := bufio.ScanWords(p, false)
	LineBreak := '\n'
	for _, pi := range p {
		if pi == byte(LineBreak) {
			wc.Line++
		}
	}
	wc.Word += advance
	return wc.Word, err
}
