package main

import "fmt"

func main() {
	a, b := calc(100) //因为calc会返回2个匿名函数，因此用两个变量a，b来装
	//此刻base的值100已经传入calc，在其作用域内生效，而calc的作用域是贯穿了add和sub的
	res1 := a(10)
	res2 := b(13)
	fmt.Println(res1)//输出值110
	fmt.Println(res2)//输出值110-13=97
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base

	}
	return add, sub

}
