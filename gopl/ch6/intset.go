package ch6

import (
	"bytes"
	"fmt"
)

//一个布隆过滤器

type IntSet struct {
	/*
		使用无符号slice
		每一个元素的每一位都表示集合里的一个值
		当集合的第i位被设置时，我们才说这个集合包含元素i
	*/
	words []uint64    
}

/*
	返回集合内的x位是否为存在
*/
func (s *IntSet)Has(x int) bool  {
	/*
		x处于第word个元素的第bit位
	*/
	word,bit := x/64,uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit)!=0
}

/*
	把一个非负数x添加到集合中（把X位置1）
*/
func (s *IntSet) Add(x int)  {
	word,bit := x/64,uint(x%64)
	/*
		words长度不够
	*/
	for word >= len(s.words) {
		s.words = append(s.words,0)
	}
	s.words[word]=s.words[word]|(1<<bit)
}

/*
	并集
*/
func (s *IntSet) UnionWith(t *IntSet)  {
	for i,tword:=range t.words{
		if i<len(s.words){
			s.words[i]= s.words[i]|tword
		}else{
			s.words = append(s.words,tword)
		}
	}
}

/*
	返回字符串格式
*/
func (s *IntSet)String() string  {
	var buf  bytes.Buffer
	buf.WriteByte('{')
	for i,word := range s.words{
		if word ==0 {
			continue
		}else {
			for j := 0;j<64;j++ {
				if word & (1 << uint(j))!=0 {
					if buf.Len()>len("{") {
						buf.WriteByte(' ')
					}
					fmt.Fprintf(&buf,"%d",64*i+j)
				}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

