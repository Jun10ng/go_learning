package ch6

/*
定义一个变参方法(*IntSet).AddAll(...int)，
这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。
*/

func (s *IntSet)AddAll(es ...int)  {
	for _,e := range es{
		s.Add(e)
	}
}