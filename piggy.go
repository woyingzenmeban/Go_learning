package main

import (
	"fmt"
	"math/rand"
)
//将0.05/0.10/0.25多次放入直至20，每次增加后显示总数，并进行适当的格式化
func main() {
	count := 0
	for piggyBank := 0.0; piggyBank < 20.00; {
		switch rand.Intn(3) {
		case 0:
			piggyBank += 0.05
		case 1:
			piggyBank += 0.10
		case 2:
			piggyBank += 0.25
		}
		fmt.Printf("%5.2f\n",piggyBank)
		count++
	}
	fmt.Println(count)		//输出一共一共执行多少次
}