package main

import (
	"fmt"
	"time"
)

func chanWithTicker() {
	ch := make(chan int)
	timer := time.NewTimer(4 * time.Second)
	ticker := time.NewTicker(500 * time.Microsecond)
	go func() {
		for {
			select {
			case b := <-ch:
				fmt.Println(b)
			case <-ticker.C:
				fmt.Println("ticker有500ms一次!")
			case <-timer.C:
				fmt.Println("timer to done")
				return
			default:
				fmt.Println("default")
				time.Sleep(time.Second)
			}
		}
	}()
	ch <- 1
	time.Sleep(10 * time.Second)
}

func main() {
	chanWithTicker()
}
