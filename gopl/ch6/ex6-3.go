package ch6

/*
 (*IntSet).UnionWith会用`|`操作符计算两个集合的并集，
我们再为IntSet实现另外的几个函数IntersectWith（交集：元素在A集合B集合均出现），
DifferenceWith（差集：元素出现在A集合，未出现在B集合），
SymmetricDifference（并差集：元素出现在A但没有出现在B，
或者出现在B没有出现在A）。
*/

//交集
func (s *IntSet)IntersectWith(t *IntSet)   {
	tlen := len(t.words)
	for i := 0;i<tlen*64;i++ {
		if t.Has(i) {
			if !s.Has(i) {
				t.Remove(i)
			}
		}
	}
}
// 第二种交集写法的效果更佳 详情看test文件的基准测试
func (s *IntSet)IntersectWith2(t *IntSet)  {
	for i,_ := range t.words{
		if i >= len(s.words){
			continue
		}else if t.words[i] == s.words[i] {
			continue
		}else {
			for j:=0;j<64;j++ {
				if (t.words[i]&(1<<uint(j)))!=0{
					if s.words[i]&(1<<uint(j))==0 {
						t.words[i]^=(1<<uint(j))
					}
				}
			}
		}
	}
}

//差集
func (s *IntSet)DifferenceWith(t *IntSet){
	for i ,sword := range s.words{
		if i>=len(t.words) {
			t.words = append(t.words, sword)
		}else if sword==t.words[i] {
			t.words[i]=0
		}else {
			for j:=0;j<64;j++ {
				if sword&(1<<uint(j))!=0 && t.words[i]&(1<<uint(j))!=0 {
					// 去掉相同数
					t.words[i]^=1<<uint(j)
				}else if sword&(1<<uint(j))!=0 && t.words[i]&(1<<uint(j))==0 {
					//s有而t没有
					t.words[i] |= (1<<uint(j))
				}
			}
		}
	}
}