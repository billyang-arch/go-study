package main

import "fmt"

type Person struct {
	Name string
	age  int
}

type Student struct {
	Person
	ClassName string
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		age:  age,
	}
}

func NewStudent(className string) *Student {
	return &Student{
		ClassName: className,
	}
}

func main() {
	s := NewStudent("123")
	fmt.Println(s)
	fmt.Println(s.Name)
}
