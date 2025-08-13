package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := make([]int, 5) //这时候创建了一个长度为5的切片，cap也是5
	fmt.Println(s, len(s), cap(s))
	PrintSliceStruct(&s)
	test(s)
}

func test(s []int) {
	PrintSliceStruct(&s)
}

func PrintSliceStruct(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))

	// 查看slice的结构
	fmt.Printf("slice struct: %+v,slice is %v\n", ss, s)
}
