package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	err := fetchUrlDemo2()
	if err != nil {
		fmt.Println("fetchUrlDemo2 failed:", err)
	}
}

// fetchUrlDemo2 使用errgroup并发获取url内容
func fetchUrlDemo2() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}
	for _, url := range urls {
		//如果不赋值一个，go func（）协程里面读取 url的值的时候，已经变成新的值了，因为for循环读取出来的值都是放在一个内存地址中的。
		//赋值给新的变量读取的时候就近原则，就会去新的变量里面去拿，不这么写也行的，可以用 go func（x string）{...}（url），直接在开启一个协程的时候把url当做参数传递过去
		url := url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	if err := g.Wait(); err != nil {
		// 处理可能出现的错误
		fmt.Println(err)
		return err
	}
	fmt.Println("所有goroutine均成功")
	return nil
}
