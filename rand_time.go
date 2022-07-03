package main

import (
	"fmt"
	"math/rand"
)

func main() {
//随机展示10次年月日，要求闰年 2月29天，平年2月28天
	for count :=0;count < 10;count ++ {
	year := rand.Intn(9999) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31

	leap := year%400 == 0 || (year%4 ==0 && year%100 != 0)

/* 	if leap  {
		switch month {
		case 2:
			daysInMonth = 29
		case 4,6,8,10:
			daysInMonth = 30
		} 
		}else {
			switch month {
			case 2:
				daysInMonth = 28
			case 4,6,8,10:
				daysInMonth = 30			
			}
		}	 */	
		switch month {
		case 2:
			daysInMonth = 28
			if leap {
				daysInMonth = 29
			}
		case 4,6,9,11:
			daysInMonth = 30
		}
		
		day := rand.Intn(daysInMonth) + 1
		fmt.Printf("Rand year %d;	Rand month %d;	Rand day %d\n", year, month, day)
	}
 }