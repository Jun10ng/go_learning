package ch6

func (s *IntSet)Len()int  {
	ans := 0
	for _ , word := range s.words{
		if word == 0 {
			continue
		}else {
			for j := 0;j < 64;j++ {
				if word & (1 << uint(j))!=0 {
					//fmt.Printf("word & (1 << uint(%d)):%v \n", j,word & (1 << uint(j)))
					ans+=1
				}
			}
		}
	}
	return ans
}

func (s *IntSet)Remove(x int)  {
	word, bit := x/64, x%64
	s.words[word] ^= (1<<uint(bit))
}

func (s *IntSet)Clear()  {
	for i,_ := range s.words{
		s.words[i] = uint64(0)
	}
}

func (s *IntSet) Copy() *IntSet  {
	cpy := &IntSet{
		words:make([]uint64,s.Len()) }
	copy(cpy.words,s.words)
	return cpy
}