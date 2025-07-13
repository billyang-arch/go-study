package main

import "fmt"

// 5,因为返回值没有变量名，所以返回值是拷贝了x的值，x的变化与返回值无关
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 6，返回值为x,相当于把5赋值给x，然后x++
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 5,返回值为y，相当于x的值拷贝给y，然后x++,y的值并不会有变化
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 5，因为defer中x的形参，相当于吧返回值x的值拷贝给了defer中x的形参，所以defer中x++，不会影响返回值x的值
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// defer语句的执行时机：返回值赋值->defer语句->返回
// 多个的defer语句的执行顺序：后进先出
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
