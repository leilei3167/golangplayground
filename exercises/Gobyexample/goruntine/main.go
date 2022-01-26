package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {

		fmt.Println(from, ":", i)
	}

}

func main() {
	f("motherfucker")
	go f("协程开启")

	go func(a string) {
		fmt.Println(a)
	}("msg")

	time.Sleep(time.Second * 2)
}
