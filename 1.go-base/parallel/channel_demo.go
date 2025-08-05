package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f3(ch)

	//对空的channel进行接收和发送操作会阻塞当前goroutine
	//关闭nil的channel会panic
	var ch1 chan int
	ch1 <- 1
}

// channel closed
// 会先获取完channel中的值，然后再判断channel是否关闭，如果channel关闭，则ok为false，否则ok为true。
//
//v:1 ok:true
//v:2 ok:true
func f2(ch chan int) {
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			return
		}
		fmt.Printf("v:%#v ok:%#v\n", val, ok)
	}
}

// 通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。上面那个示例我们使用for range改写后会很简洁。
func f3(ch chan int) {
	for v := range ch {
		fmt.Printf("v:%#v\n", v)
	}
}
