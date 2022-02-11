package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {

			fmt.Println("子发送:", i)
			ch <- i
			time.Sleep(time.Second)
		}

	}()
	for i := 0; i < 5; i++ {
		fmt.Println("收到:", <-ch)
	}

}
