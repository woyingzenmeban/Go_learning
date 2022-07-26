package main

import (
	"fmt"
	"log"
)

func main() {
	var  nobody *person
	fmt.Println(nobody)
	nobody.brithday()
}

type person struct{
	age  int

}

func  (p *person) brithday() {
	if p == nil {
		log.Fatalln("p is null")
		// return 
	}
	p.age++
}
