package main

import "fmt"

func main() {
	//日常指向指针及修改指针值的变化
	/* 	hello := 42
	   	fmt.Println(hello)
	   	fmt.Println(&hello)
	   	fmt.Println(*&hello)

	   	shi := &hello
	   	fmt.Println(*shi)
	   	*shi++
	   	fmt.Println(hello)		//指向原变量信息，非副本，指针内容改变，源地址也会改变
	   	fmt.Println(*shi)
	   	hello +=1
	   	fmt.Println(hello)		//指向原变量信息，非副本，源地址改变，，指针值也会改变
	   	fmt.Println(*shi)

	   	var address *int
	   	fmt.Printf("address type is %T\n", address) */

	/* //指向结构体的指针

	type person struct{
		name, superpower string
		age int
	}

	shi := &person{
		name: "shixiaoshan",
		age: 25,
	}
	shi.superpower = "China"		//下面的两种方法都可以直接修改指针所对应的结构体字段的值
	(*shi).superpower = "China"		//括号内的* 非必须写
	fmt.Printf("%+v\n", shi) */

	/* //	指向数组的指针

	superpower := &[3]string{"shi", "xiao", "shan"}
	fmt.Println(superpower[0])
	fmt.Println(superpower[len(superpower)-1])		//获取最后一位的内容
	fmt.Println(superpower[2:])
	fmt.Println((*superpower)[2:])
	*/

	//配合函数使用

	/* 	rebecca := person{name: "shi",superpower: "xiao", age: 25}
		birthday(&rebecca)

	}

	type person struct{
		name, superpower string
		age int
	}

	func birthday(p *person){
		p.age++
		fmt.Printf("%+v\n",p) */

	/* //结合方法使用

	   	shi := &person{"shi","xiao",25}
	   	shi.birthday()
	   //}

	   }
	   type person struct{
	   	name, superpower string
	   	age int
	   }

	   func (p *person) birthday(){
	   	p.age++
	   	fmt.Printf("%+v\n",p)
	   } */

	/* //修改元素值
	   	var board [8][8]rune
	   	reset(&board)
	   	fmt.Printf("%c",board[0][0])

	   }

	   func reset(board *[8][8]rune) {
	   	board[0][0] = 'r' */
	/*
	   //修改数组元素值
	   	c := []string{"s","h","i","x","i"}
	   	a_name(&c)
	   	fmt.Println(c)

	   }

	   func a_name(b *[]string) {
	   	*b = (*b)[0:2]
	   }
	*/


//练习题
	var t turtle
	t.up()
	t.up()
	fmt.Println(t)
	t.right()
	t.right()
	fmt.Println(t)

}

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y--
}
func (t *turtle) down() {
	t.y++
}
func (t *turtle) left() {
	t.x--
}
func (t *turtle) right() {
	t.x++
}
