package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancle := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got to stop chan")
				return
			default:
				fmt.Println("waiting signal!")
				time.Sleep(1 * time.Second)
			}

		}

	}()

	time.Sleep(time.Second * 5)
	fmt.Println("i want to stop")
	cancle()
	time.Sleep(time.Second * 2)
}
