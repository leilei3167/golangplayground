package main

import (
	"fmt"
	"runtime"
)

/* 默认情况下通道是无缓冲的,意味着只有对应的接收<-chan准备好了接收时
才允许发送,可缓存通道允许在没有对应接收方的情况下,缓存限定数量的值 */

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)//一百个缓存空间的通道

	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i

		}
		close(ci)
		c <- struct{}{}

	}(c, ci)

	fmt.Println("当前协程数量:", runtime.NumGoroutine())

	//主协程不断取数据,来实现同步等待

	<-c

	//此刻ci已经关闭,协程已退出,但可以继续取数据
	fmt.Println("当前协程数量:", runtime.NumGoroutine())

	for v := range ci {
		fmt.Println(v)

	}

}
