package main

import "fmt"

type Child struct {
	age int
}

func (s *Child) addAge() {
	s.age += 1
}

func (s Child) addAge1() {
	s.age += 1
}

func main() {

	child := &Child{age: 1}
	child.addAge()
	fmt.Println(child.age)
	child.addAge1()
	fmt.Println(child.age)

}
