package main

import (
	"fmt"
	"time"
)

/* default 语句是可选的；fallthrough 行为，和普通的 switch 相似，是不允许的。
在任何一个 case 中执行 break 或者 return，select 就结束了。

select 做的就是：选择处理列出的多个通信情况中的一个。

如果都阻塞了，会等待直到其中一个可以处理
如果多个可以处理，随机选择一个
如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。
在 select 中使用发送操作并且有 default 可以确保发送不被阻塞！如果没有 case，select 就会一直阻塞。

select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出。

*/
func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
		time.Sleep(time.Second)
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
		time.Sleep(time.Second * 2)

	}
}

func suck(ch1, ch2, ch3 chan int) { //无限循环select,对应的case有输入或输出即执行
	for {
		select {
		case v := <-ch1:
			fmt.Printf("recevied on channel1 :%d\n", v)
			ch3 <- 1
		case v := <-ch2:
			fmt.Printf("recevied on channel2 :%d\n", v)
			ch3 <- 1
		}

	}

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int, 10)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2, ch3)

	for {
		if len(ch3) == 10 {
			break
		}
	}
	fmt.Println("执行完毕")

}
