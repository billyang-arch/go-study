package main

import "fmt"

// Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，
// 例如： 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为false。 切片、函数、指针变量的默认为nil。
func main() {
	var m int
	fmt.Println(m)

	//标准格式
	var a int = 1
	fmt.Println(a)

	var b = "2"
	fmt.Println(b)

	//只能在函数内使用
	c := 3
	fmt.Println(c)

	d, e := 4, 5
	fmt.Println(d, e)

	//批量声明
	//var (
	//	a1 string
	//	b1 int
	//	c1 bool
	//	d1 float32
	//)
}
