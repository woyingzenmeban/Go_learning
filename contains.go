package main

import (
	"fmt"
	"strings"
)

func main() {
	command := "walk outside"
	exit := strings.Contains(command, "walk")
	fmt.Println("Yes", exit)	//返回True
}