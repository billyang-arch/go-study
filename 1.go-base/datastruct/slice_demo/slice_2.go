package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := make([]int, 5)

	case1(s)
	case2(s)
	case3(s)
	case4(s)
	printSliceStruct(&s)
}

func case1(s []int) {
	s = s[1:]
	s[0] = 1
	printSliceStruct(&s)
}
func case2(s []int) {
	s = s[1:3]
	printSliceStruct(&s)
}
func case3(s []int) {
	s = s[len(s)-1:]
	printSliceStruct(&s)
}
func case4(s []int) {
	s = s[2:]
	printSliceStruct(&s)
}
func printSliceStruct(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))

	// 查看slice的结构
	fmt.Printf("slice struct: %+v,slice is %v\n", ss, s)
}
