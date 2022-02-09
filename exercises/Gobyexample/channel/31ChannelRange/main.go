package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int, 5)

	go func() {
		for v := range chan1 {
			fmt.Println(v)
		}
	}()

	for i := 1; i < 6; i++ {
		chan1 <- i
		time.Sleep(time.Second)
	}
	close(chan1) //一个非空的通道也是可以关闭的， 但是通道中剩下的值仍然可以被接收到。

	for v := range chan1 {
		fmt.Println(v)
	}

}
