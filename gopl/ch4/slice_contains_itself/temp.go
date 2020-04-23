package slice_contains_itself

import "fmt"

/*
gopl 4.2章中说：
	it possible for a slice to contain itself.
	切片可能包含它本身。
探讨下如何报包含
*/

func sliceContainsItself()  {
/*	//证明一
	//这段代码运行会报 stack overflow
	//因为是s[2]是个自身，不断取值，不断循环
	//这也证明了slice是可以指向自身的，
	//但是不能被打印出来
	s := []interface{}{"one",2,nil}
	s[2] = s
	fmt.Printf("s is %v\n",s)
	fmt.Printf("s[2] is %v",s[2])
*/


/*	//证明二
	//这个例子说明了 s2为s的第三个元素，
	//s2不但是个切片而且，前两个值和s相同
	//避开了会导致溢出的第三个元素
	//用另一种方式证明了切片能包含它自身
	s:=[]interface{}{"one","two",nil}
	s[2] = s
	s2 := s[2].([]interface{})
	fmt.Printf("s2[0] is %v, s2[1] is %v \n",s2[0],s2[1])
	s3 := s2[2].([]interface{})
	fmt.Printf("s3[0] is %v, s3[1] is %v \n",s3[0],s2[1])*/

	//s数组为什么不能包含自身
	//当s被放入[]inteface{}时，会创建一个副本
	//所以已经不是自身了
	s := [2]interface{}{"one", nil}
	s[1] = s
	fmt.Printf("s is %v \n",s) //output:[one [one <nil>]]
	fmt.Printf("s address is %p ,while s[1] address is %p.\n",&s,&s[1])
	// output: s address is 0xc000004560 ,while s[1] address is 0xc000004570.
	//明显看得出两个地址并不一样
	fmt.Println(s[0])
	s2 := s[1].([2]interface{})
	fmt.Println(s2[0])
	s3 := s2[1].([2]interface{})
	fmt.Println(s3)


}