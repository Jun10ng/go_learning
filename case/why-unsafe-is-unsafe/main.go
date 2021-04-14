package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {

	a := "aaa"
	now := time.Now()
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Println(time.Since(now))
	fmt.Printf("%T", b)
	fmt.Printf("%v", b)
}
