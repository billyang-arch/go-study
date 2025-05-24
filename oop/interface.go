package main

import "fmt"

type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

func (l *Logger) Log(msg string) {
	println(msg)
}

func (l *Logger) Start() {
	println("start")
}

func NewLogger() *Logger {
	return &Logger{}
}

func testService(s Service) {
	s.Log("test")
	s.Start()
}

type CustomLogger struct {
	Logger
}

func (c *CustomLogger) Start() {
	println("custom start")
}

func main() {
	//s := NewLogger()
	s1 := &CustomLogger{}
	testService(s1)

	var any interface{}

	any = 1
	fmt.Println(any)
	fmt.Printf("%T\n", any)

	any = "hello"
	fmt.Println(any)
	fmt.Printf("%T\n", any)

}
