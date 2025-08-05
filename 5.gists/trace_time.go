package main

import (
	"fmt"
	"time"
)

func main() {
	defer TrackTime(time.Now())

	time.Sleep(500 * time.Millisecond)
}

func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed:", elapsed)

	return elapsed
}

//elapsed: 501.134583ms
