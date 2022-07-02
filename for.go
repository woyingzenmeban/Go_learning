package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10

	// for count > 0 {			#与下方结果一样
	for {
		if count < 0 {
			break
		}
		fmt.Println(count)
		time.Sleep(time.Second * 2)		//等待2s
		count--
	}
	fmt.Println("liftoff")
}