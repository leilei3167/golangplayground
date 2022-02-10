package main

import (
	"fmt"
	"math/rand"
)

//多个协程生成器,利用select的扇入技术.Fan in
//用select聚合多条通道服务

func GenerateIntA() chan int {

	ch := make(chan int, 10)
	//启动协程来生成随机数
	go func() {
		for {
			ch <- rand.Int()
		}

	}()
	return ch

}
func GenerateIntB() chan int {

	ch := make(chan int, 10)
	//启动协程来生成随机数
	go func() {
		for {
			ch <- rand.Int()
		}

	}()
	return ch

}

func GenerateInt() chan int {
	ch := make(chan int, 20)
	go func() {

		for {
			select {
			case ch <- <-GenerateIntA():

			case ch <- <-GenerateIntB():

			}

		}

	}()
	return ch
}

func main() {
	ch := GenerateInt()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)

	}

}
