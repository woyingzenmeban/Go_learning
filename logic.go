package main

import "fmt"

func main() {
	fmt.Println("This is logic compare file.Test current year is leap year?")
	year := 2100
	leap := year%400 == 0 || (year%4 ==0 && year%100 != 0)		//判断闰年的方法
	if leap {			//leap变量为真，即是闰年
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not a leap year")
	}

}