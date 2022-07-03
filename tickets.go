package main

import (
	"fmt"
	"math/rand"
)

/*随机生成4列表格，包含
Spacelne:航空公司，xxxxxxxx，参考下方
Days：到火星的单程天数，距离62100000公里，每秒随机16-30km/s
Trip type：单程或往返
Price：百万美元，随机生成在3600万-5000万之间（单程）

*/
const secondsPerday = 86400

func main() {
	company := ""
	Trip := ""
	distance := 62100000
	fmt.Println("Spaceline           Days  Trip type  Price")
	fmt.Println("==========================================")

	for count:= 0;count  < 10;count ++ {
	switch rand.Intn(3) {
	case 0:
		company = "Space Adventures"
	case 1:
		company = "SpaceX"
	case 2:
		company = "Virgin Galactic"
	}
	speed := rand.Intn(15) + 16			//速度
	day := distance / speed / secondsPerday		//所需天数
	price := 20 + speed		//价格
	if rand.Intn(2) == 1 {
		Trip = "Round-trip"
		price = price * 2
	} else {
		Trip = "One=way"
	}

	fmt.Printf("%-17v %6v %11v %5v\n",company, day, Trip, price)
	}
}