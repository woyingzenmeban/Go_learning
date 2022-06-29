package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "shi xiao shan and you"
	if name == "hello " {
		fmt.Println("you guess error")
	} else if name == "Hello" {
		fmt.Println("you guess error again")	
	} else if name == "shan" {				// 只写部分内容不会进行匹配
			fmt.Println("you guess little right")
	} else if name == "shi xiao shan and you" {				// 这里匹配必须完全匹配才能识别
		fmt.Println("you guess all right")
	} else if strings.Contains(name, "shan") {				//若没有上面内容，则会匹配这行
		fmt.Println("you guess good")
	} else {
		fmt.Println("Don't guess")
	}

}