package ex4_6

import (
	"unicode"
)

func replaceSpace(b []byte) []byte {
	end := len(b)
	for i := 0; i < len(b); i++ {

		if !unicode.IsSpace(rune(b[i])) {
			continue
		} else {
			spaceCount := 0
			cpyStart := i
			for unicode.IsSpace(rune(b[i])) {
				i++
				spaceCount++
			}
			end = end - spaceCount + 1
			copy(b[cpyStart:], b[cpyStart+spaceCount-1:])
			i = cpyStart
		}
	}
	return b[:end]
}
