package main

import (
	"fmt"
	"time"
)

func fun1(a, b int) (int, int, int) { //未命名的返回值
	return a - b, a + b, a * b
}
func fun2(a, b int) (x1 int, x2 int, x3 int) { //命名的返回值，在实际中应尽量采取这种方式，提高可读性
	x1 = a - b
	x2 = a + b
	x3 = a * b
	return
}

//写一个函数，使其能接受一个变长参数并对其元素进行打印
func fun3(s ...int) {
	for k, v := range s {
		fmt.Println(k, v)
	}
}

func main() {
	start := time.Now()
	a, b, _ := fun2(1, 2) //对于有多个返回值的情况，可以用下划线（空白符）来接收不需要的返回值
	end := time.Now()
	delta := end.Sub(start)//统计这个函数的执行时间
	fmt.Println(delta)
	fmt.Println(a, b)
	c := fun1 //有函数名的函数，可以作为赋值给一个变量，后续调用函数可以调用被赋值的变量，如下
	d, _, _ := c(1, 4)
	fmt.Println(d)
	e := []int{1, 2, 3, 4, 5}
	fun3(e...)

}
