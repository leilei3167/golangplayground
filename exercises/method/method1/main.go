package main

import "fmt"

/* 一般结构体包含各类数据，通过方法来给结构体定义行为
 */
type Person struct { //定义一个结构体
	Name string
	Sex  string
	Age  int
}

func (p1 Person) speak() { //定义一个只有Person类型才能调用的方法，方法名字为speak，此处p1非常类似于传参
	fmt.Println(p1.Name, "是一个好人")
}

func (p2 Person) jisuan() { //可以定义一个方法来计算
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(res)

}

func (p3 Person) jisuan2(n int) { //定义一个有参数输入的方法
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(res)

}

func (p4 Person) getsum(n1 int, n2 int) int { //可以定义有参数和返回值的方法
	return n1 + n2

}

func main() {
	jack := Person{Name: "jack", Sex: "male", Age: 67} //定义一个结构体变量jack

	jack.speak()               //调用方法speak
	jack.jisuan()              //调用计算
	jack.jisuan2(10)           //传参调用jisuan2
	res := jack.getsum(10, 20) //因为有返回值，需要一个变量来接收，传递10和20 的到30返回值
	fmt.Println(res)
}
