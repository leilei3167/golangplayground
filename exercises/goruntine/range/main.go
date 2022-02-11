package main

import "fmt"

func main() {
	n := make(chan int)
	m := make(chan int)

	go func() {

		for x := 0; x < 100; x++ {
			n <- x

		}
		close(n)
	}()

	go func() {

		for i := 0; i < 100; i++ {
			m <- <-n * 5

		}
		close(m)
	}()

	for y := range m {
		fmt.Println(y)

	}

}
