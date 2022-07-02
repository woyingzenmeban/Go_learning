package main

import (
	"fmt"
)

func main() {
	fmt.Println("There is a cavern entrance here and a path to the east")

	command := "go inside"

	switch command {
	case "go east":
		fmt.Println("you head further up the mountain")
	case "enter cave", "go inside":				//必须完全匹配变量内容
		fmt.Println("you find yourself a adimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'")
	default:
		fmt.Println("Didn't quite get that")
	}

		
}