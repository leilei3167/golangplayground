package main

import (
	"fmt"
)

func main() {
	message := make(chan string)
	signal := make(chan bool)

	//非阻塞接收,如果message中存在值,进入msg分支,否则执行default
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")

	}

	//非阻塞发送

	msg := "hi"
	select {

	case message <- msg:
		fmt.Println("snet mes:", msg)
	default:
		fmt.Println("no mes sent")

	}

	//可以在default前使用多个case实现一个多路的非阻塞选择器

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signal:

		fmt.Println("received signal", sig)
	default:
		fmt.Println("nonono")
	}

}
