package main

import (
	"fmt"
)

/* type report struct {
	sol		int			//日期
	high, low	float64		//温度:最高温和最低温
	lat, long	float64		//坐标，经纬度
} */


/* type report struct {
	sol		int			//日期
	temperature temperature //字段名，下方的结构体类型	
	location location  //字段名，下方的结构体类型
} */

type sol int

type report struct {
	sol					//日期
	temperature  //下方的结构体类型	，无字段名，但一样可以在顶级struct组合中使用
	location   //下方的结构体类型，无字段名，
}


type  temperature  struct {
	high, low	celsius		//温度:最高温和最低温；类型又引入了最下方的结构体
}

type location struct {
	lat, long	float64		//坐标，经纬度
}

type celsius float64


func (t temperature) average() celsius {			//添加一个average的平均方法，类似于多了一个average的字段
	return (t.high + t.low) / 2
}

/* func (r report) average() celsius {				//根据report结构体，返回温度字段的方法-平均值
	return r.temperature.average()
} */


func (s sol) days(s2 sol) int {
	return 11
}

func (l location) days(l2 location) int {
	return 22
}

/* func (re report） days(s2 sol) int{
	return re.sol.days(s2)
} */

func main() {
	bradbury := location{-412.25, 153545}
	t := temperature{high: -1.0, low: -78.0}
	fmt.Println(t.average())		//使用上方的temperature 结构体及函数，返回average的值

	report := report{
		sol: 15,
		temperature: t,
		location: bradbury,
	}
	fmt.Println(report.average())

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v℃ \n", report.high)
	fmt.Println(report.sol.days(1111))
	// fmt.Println(report.days(11111))			//此时就会命名冲突
}