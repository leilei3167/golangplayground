package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

//
func produce(done chan struct{}) chan int {
	ch := make(chan int) //无缓冲

	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done: //监听是否发出退出通知信号
				break Lable
			}
		}
		//收到通知后关闭ch
		close(ch)
	}()
	return ch

}

func main() {
	done := make(chan struct{})

	ch := produce(done)

	fmt.Println(<-ch) //生成1个
	fmt.Println(<-ch)
	fmt.Println(runtime.NumGoroutine())
	done <- struct{}{} //发送通知关闭
	//func close(c chan<- Type)
	v, ok := <-ch
	fmt.Println(v, ok)
	fmt.Println(<-ch) //再取就是0值
	fmt.Println(<-ch)

}
