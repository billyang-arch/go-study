package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func createNum(ch chan int64) {
	for i := 0; i < 50; i++ {
		random := rand.Int63n(100000)
		fmt.Println("随机数:", random)
		ch <- random
	}
	close(ch)
	fmt.Println("关闭管道")
}

func calNumSum(jobChan chan int64, resultChan chan int64) {
	for v := range jobChan {
		var res int64
		for v > 0 {
			res += v % 10
			v /= 10
		}
		resultChan <- res
	}
	wg1.Done()
}

var wg1 sync.WaitGroup

func main() {
	jobChan := make(chan int64, 100)
	resultChan := make(chan int64, 100) //这里缓冲区必须大于随机数的数量，否则会出现死锁
	go createNum(jobChan)
	for i := 0; i < 24; i++ {
		wg1.Add(1)
		go calNumSum(jobChan, resultChan)
	}

	wg1.Wait()
	close(resultChan)
	//管道必须关闭，否则这里的for循环会一直阻塞
	for v := range resultChan {
		fmt.Println("结果:", v)
	}
}
