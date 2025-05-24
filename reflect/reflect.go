package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"lisi", 13}

	fmt.Println(reflect.ValueOf(p))        // {lisi 13}  变量的值
	fmt.Println(reflect.ValueOf(p).Type()) // main.Person 变量类型的对象名

	fmt.Println(reflect.TypeOf(p)) //  main.Person	变量类型的对象名

	fmt.Println(reflect.TypeOf(p).Name()) // Person:变量类型对象的类型名
	fmt.Println(reflect.TypeOf(p).Kind()) // struct:变量类型对象的种类名

	fmt.Println(reflect.TypeOf(p).Name() == "Person")       // true
	fmt.Println(reflect.TypeOf(p).Kind() == reflect.Struct) //true

	//反射操作简单数据类型
	var num int64 = 100

	// 设置值：指针传递
	//这里传递的是指针而不是值本身，因为后续需要通过反射来修改这个值
	//在Go反射中，如果要修改一个值，必须传递它的指针而不是值本身
	//获取指针的反射对象后，可以通过Elem()方法获取指针指向的实际值的反射对象(如后续代码所示)
	ptrValue := reflect.ValueOf(&num)
	fmt.Printf("ptrVal %T  %v\n", ptrValue, ptrValue) // *int64
	newValue := ptrValue.Elem()                       // Elem()用于获取原始值的反射对象
	fmt.Println("type：", newValue.Type())             // int64
	fmt.Println(" can set：", newValue.CanSet())       // true
	newValue.SetInt(200)

	// 获取值：值传递
	rValue := reflect.ValueOf(num)
	fmt.Println(rValue.Int())               // 方式一：200
	fmt.Println(rValue.Interface().(int64)) // 方式二：200

	//反射进行类型推断
	type user struct {
		Name string
		Age  int
	}

	u := &user{
		Name: "Ruyue",
		Age:  100,
	}

	fmt.Println(reflect.TypeOf(u))         // *main.user
	fmt.Println(reflect.TypeOf(*u))        // main.user
	fmt.Println(reflect.TypeOf(*u).Name()) // user
	fmt.Println(reflect.TypeOf(*u).Kind()) // struct
}
