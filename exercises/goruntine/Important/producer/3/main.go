package main

import (
	"fmt"
	"math/rand"
)

//退出通知机制实现自动退出
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int, 10)

	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done: //当done能被取时,说明收到信号退出
				break Lable
			}

		}
		close(ch)
	}()
	return ch
}

func main() {
	//信号
	done := make(chan struct{})
	ch := GenerateIntA(done) //传入一个信号通道

	for i := 0; i < 5; i++ {

		fmt.Println(<-ch)
	}
	//取值完成后通知生成器关闭
	done <- struct{}{}
	comma, ok := <-ch
	fmt.Println(comma, ok)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

}
