package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
	wg.Done()
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	fmt.Println("你好")
	wg.Wait()
}

//多次执行上面的代码会发现每次终端上打印数字的顺序都不一致。这是因为10个 goroutine 是并发执行的，而 goroutine 的调度是随机的。

// goroutine不同于线程，他运行在用户态，而线程是运行在内核态。
// goroutine的调度是go语言的运行时自己实现的，而线程的调度是由操作系统来实现的,m个协程对应n个线程，有go语言自动分配到具体的线程上，相当于自动实现了线程池。
// goroutine的栈空间很小，只有2KB，而线程的栈空间一般至少1-2MB。所以理论上可以创建成百上千上万个goroutine，而操作系统对线程数是有限制的。
// go语言通过通信共享内存，而不是通过共享内存来通信。
