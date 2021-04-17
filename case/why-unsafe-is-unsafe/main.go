package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//go:noinline
func StringToByte(k *string) []byte {
	ptr := (*reflect.SliceHeader)(unsafe.Pointer(k))
	ptr.Cap = ptr.Len
	return *(*[]byte)(unsafe.Pointer(ptr))
}

//func main() {
//
//	x := "12345678990"
//	a := x[0:6]
//	b := x[2:8]
//	//ssa := *(*reflect.SliceHeader)(unsafe.Pointer(&a))
//	//ssb := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
//	//aa := *(*[]byte)(unsafe.Pointer(&ssa))
//	//bb := *(*[]byte)(unsafe.Pointer(&ssb))
//	//fmt.Println(string(aa))
//	//fmt.Println(string(bb))
//	//
//	fmt.Println(string(StringToByte(&a)))
//	fmt.Println("=======")
//	fmt.Println(string(StringToByte(&b)))
//
//}

//func StringToByte(key *string) []byte {
//	strPtr := (*reflect.SliceHeader)(unsafe.Pointer(key))
//	strPtr.Cap = strPtr.Len
//	b := *(*[]byte)(unsafe.Pointer(strPtr))
//	return b
//}

func main() {
	decryptContent := "/AvYEjm4g6xJ3LVrk2/Adksdfsdfssssssss"
	iv := decryptContent[0:16]
	key := decryptContent[23:24]

	//fmt.Println(&iv)
	//fmt.Println(&key)
	ivBytes := StringToByte(&iv)
	keyBytes := StringToByte(&key)
	fmt.Println(string(ivBytes))
	fmt.Println(string(keyBytes))
}
