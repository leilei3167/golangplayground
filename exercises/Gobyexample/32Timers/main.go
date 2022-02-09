package main

import (
	"fmt"
	"time"
)

//定时器

func main() {
	//创建一个定时2秒的定时器,将提供一个timer结构体,包含timer.C通道,
	//直到C发送定时器失效的值之前将一直阻塞
	timer1 := time.NewTimer(time.Second * 2)
	//
	a := <-timer1.C
	fmt.Println("Timer 1 完成")
	fmt.Println(a)

	//如果只是单纯的等待用Time.Sleep即可,定时器有用的原因之一就是他
	//可以在失效前取消

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 完成")

	}()
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("timer2 stopped")
	}

}
