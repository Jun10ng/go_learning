package ch6

import (
	"fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	var x,y  IntSet
	x.AddAll(1,9,144)
	fmt.Println(x.String()) //"{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())//"{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())//"{1 9 42 144}"
	fmt.Println(x.Has(9),x.Has(123))
	// Len
	fmt.Printf("x len : %d \n",x.Len())
	// Remove
	x.Remove(9)
	fmt.Printf("Remove element 9 : %v \n",x.String())
	// Copy
	cpy := x.Copy()
	fmt.Printf("Copy of x : %v \n",cpy.String())
	// Clear
	cpy.Clear()
	fmt.Printf("after copy clearing %v \n",cpy.String())
	//交集
	//x :{1 42 144} 	 y:{9 42}
	x.IntersectWith(&y)
	fmt.Printf("x IntersecWit y :%v\n",y.String())
	//差集
	x.DifferenceWith(&y)
	fmt.Printf("x DifferenceWith y :%v\n",y.String())
}

func BenchmarkIntSet_IntersectWith(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		s ,y := IntSet{},IntSet{}
		s.AddAll(1,2,3,64,65,66,131,138,139)
		y.AddAll(11,22,3,164,165,66,31,38,139)
		s.IntersectWith(&y)
	}
}

func BenchmarkIntSet_IntersectWith2(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		s, y := IntSet{}, IntSet{}
		s.AddAll(1, 2, 3, 64, 65, 66, 131, 138, 139)
		y.AddAll(11, 22, 3, 164, 165, 66, 31, 38, 139)
		s.IntersectWith(&y)
	}
}

