package main

import "fmt"

func main() {
	//channel的使用
	//1.创建一个可以存放3个int的管道
	var intChan chan int
	intChan = make(chan int, 3) //必须初始化才能使用，注意容量不会自动增长的

	//2.看看intChan是什么
	fmt.Printf("intChan的值%v\n", intChan)   //为一个地址
	fmt.Printf("intChan的地址%p\n", &intChan) //intChan本身的地址，这个地址存放的值也是一个地址
	//3.写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 1
	//intChan <- 11，超过容量，将会报错fatal error: all goroutines are asleep - deadlock!

	//4.看看这是候的长度和容量
	fmt.Printf("len=%v,cap=%v\n", len(intChan), cap(intChan))
	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan //先进先出，会取出第一个装入管道的数10
	fmt.Println(num2)
	fmt.Printf("len=%v,cap=%v\n", len(intChan), cap(intChan)) //取出一个数长度变成了2，但是容量不会变
	//6.在没有使用协程的情况下，如果管道中的数已被全部取出，再取就会报错
	num3 := <-intChan
	<-intChan //取数据是可以不接收的
	fmt.Println(num3)
	//	num5 := <-intChan  //报错，因为取出超过了容量
	//fmt.Println(num5)

}
