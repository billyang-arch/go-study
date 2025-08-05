package main

import (
	"fmt"
	"time"
)

func main() {
	var d = int64(0)
	defer func(d *int64) {
		fmt.Printf("& %v Unix Sec\n", *d)
	}(&d)
	fmt.Println("Done ")
	d = time.Now().Unix()
	fmt.Println("Done ", d)

}
