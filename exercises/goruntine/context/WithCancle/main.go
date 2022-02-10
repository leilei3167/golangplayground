package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) <-chan int {
	ch := make(chan int, 2)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}

		}
	}()
	return ch
}

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	for n := range worker(ctx) {
		fmt.Println(n)
		if n == 6 {
			break
		}

	}

}
