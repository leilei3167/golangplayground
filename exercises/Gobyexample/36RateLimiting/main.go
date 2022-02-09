package main

import (
	"fmt"
	"time"
)

//速率限制是一个重要的控制服务资源和质量的途径

func main() {
	//假设我们想限制我们接收请求的处理,我们将这些请求发给一个相同的通道

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i

	}
	close(requests)

	//设置速率限制的管理器,每200ms接收处理一个值
	limiter := time.Tick(time.Millisecond * 200)
	//Tick是NewTicker的封装,只提供其内部字段C 的访问

	for req := range requests {
		<-limiter //limiter每隔200ms才会有一个发送值,因此此函数段将会间隔循环
		fmt.Println("request:", req, time.Now())
	}

	//有时候我们想临时进行速率限制,并且不影响整体的速率控制,可以通过管道缓冲来实现

	fmt.Println("-------------------------------------")
	//用来进行3次临时的脉冲型速率限制
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ { //先添加3个缓冲值
		burstyLimiter <- time.Now()

	}

	//每200ms添加一个新值
	go func() {
		for t := range time.Tick(time.Millisecond * 1000) {
			burstyLimiter <- t

		}
	}()
	//模拟超过5个接入请求

	burstyRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequests <- i

	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter //因为事先有3个已存储好的值,所以前3次可以快速执行,后2次将会间隔200ms
		fmt.Println("request", req, time.Now())
	}

}
