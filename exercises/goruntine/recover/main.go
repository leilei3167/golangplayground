package main

import (
	"fmt"
	"time"
)

/* 其中一个协程出panic，但是我们不想让他导致整个程序崩溃
可以使用defer recover来解决 */
func sayHello() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("hello world")

	}
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test发生错误", err)

		}
	}() //匿名函数调用函数体后加个小括号
	var myMap map[int]string
	myMap[0] = "golang" //不初始化map，故意犯错,panic: assignment to entry in nil map

}

func main() {
	go sayHello()
	go test()
	time.Sleep(time.Second * 6)
}
