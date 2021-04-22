package opslice

// GOSSAFUNC=newSlice go build op-slice.go
func newSlice() []int64 {
	arr := [6]int64{1, 2, 3, 4,5,8}
	slice := arr[1:3]
	return slice
}

