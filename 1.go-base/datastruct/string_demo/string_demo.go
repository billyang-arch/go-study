package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	a := []byte{1, 2, 3}
	b := strings.Builder{}
	b.Write(a)
	b2 := bytes.NewBuffer(a)

	str1 := b.String()
	str2 := b.String()
	fmt.Println(str1, str2)
	String2Bytes(str1)
	String2Bytes(str2)

	str3 := b2.String()
	str4 := b2.String()
	fmt.Println(str3, str4)
	String2Bytes(str3)
	String2Bytes(str4)
}

func String2Bytes(s string) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}

	fmt.Println(bh.Data)
}
