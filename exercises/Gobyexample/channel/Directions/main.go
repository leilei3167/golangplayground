package main

import "fmt"

/* 当使用通道作为函数的参数时,你可以指定这个通道是用来接收还是发送
这个特性提高了程序类型的安全性 */

func ping(pings chan<- string, msg string) {
	//参数是一个只允许发送数据的channel\
	pings <- msg

}

func pong(pings <-chan string, pongs chan<- string) {
	//只允许接收来自通道的数据
	msg := <-pings //pings是只能向管道放入数据的类型
	pongs <- msg   //pongs只能是只能接收通道里的值的类型
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
