package main

import (
	"fmt"
	"time"
)

/* 超时 对于一个连接外部资源，或者其它一些需要花费执行时间 的操作的程序而言是很重要的。
得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。 */

func main() {
	//假设执行一个外部调用
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "res 1"

	}()

	//使用select实现一个超时操作
	//After设置为1秒,超过1秒的将不会执行
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		//After会在另一线程经过时间段d后向返回值发送当时的时间。等价于NewTimer(d).C。
		fmt.Println("timeout 1")

	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "res 2"

	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")

	}

}
