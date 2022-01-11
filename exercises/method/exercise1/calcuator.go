package main

import "fmt"

type number int //定义一个自定义类型

func (a number) print(n int) { //给该自定义类型创建一个方法来打印乘法表
	if n < 1 {
		fmt.Println("输入的值有误，请重新输入")
		fmt.Scanf("%v", &n)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v\t", j, i, i*j)
		}
		fmt.Println()
	}

}

func main() {
	var x number
	fmt.Print("请输入1-9的值")
	fmt.Scanf("%d", &x)
	x.print(int(x))

}
