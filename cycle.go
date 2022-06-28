package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 25

	for {
		var randnum = rand.Intn(100) +1		//随机数，从1-100
		fmt.Println(randnum)
		if randnum > num {
			fmt.Println("you guess error ,more than ")
			time.Sleep(time.Second)		//等待1s
		} else if randnum < num {
			fmt.Println("you guess erro , less than")
			time.Sleep(2 * time.Second)		//等待2s
		} else {
			fmt.Println("you guess good")			
			break						//相等则退出循环
			}	
		}

}