package ex5_19

func f() (ans int){
	defer func() {
		if p:=recover();p!=nil{
			ans = p.(int)
		}
	}()
	panic(3)
}