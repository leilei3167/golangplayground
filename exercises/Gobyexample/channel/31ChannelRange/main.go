package main

import "fmt"

func main() {

	chan1 := make(chan int, 5)
	for i := 1; i < 6; i++ {
		chan1 <- i

	}
	close(chan1) //一个非空的通道也是可以关闭的， 但是通道中剩下的值仍然可以被接收到。

	for v := range chan1 {
		fmt.Println(v)
	}

}
