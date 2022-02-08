package main

import (
	"fmt"
	"runtime"
	"time"
)

//Go的选择器可以同时等待多个通道操作,Go协程和通道以及选择器的结合是Go的强大特性

func main() {
	//我们将从一下两个通道中选择
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"

	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"

	}()
	//使用select关键字来同时等待这两个值
	//当监听的通道无法读或写时,会一直阻塞,有一个可读或可写,进入该分支处理流程
	//多个可读写时则随机选取一个读取
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0)) //0表示查询当前可并发执行的协程数目,大于一表示设置
	for i := 0; i < 2; i++ {

		select {
		case msg1 := <-c1:
			fmt.Println("recevied", msg1)
		case msg2 := <-c2:
			fmt.Println("recevied", msg2)

		}

	}

}
