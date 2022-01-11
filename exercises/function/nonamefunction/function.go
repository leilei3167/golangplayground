package main

import "fmt"

func main() {

	res1 := func(n1, n2 int) int { //调用方式1，定义的时候直接调用，函数体之后直接跟实参
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)

	a := func(n1, n2 int) int { //调用方法2，将函数赋值给一个变量a
		return n1 - n2
	}
	res2 := a(11, 23)//a就成了函数变量，可以通过调用a来多次使用匿名函数
	fmt.Println(res2)

}
