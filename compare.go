package main

import "fmt"

func main() {
	fmt.Println("this is compare age")
	shi_age := 25 
	shan_age := shi_age < 18
	fmt.Printf("shi age is %v, shan age is minor? %v\n", shi_age, shan_age)
}