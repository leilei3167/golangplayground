package main

import (
	"fmt"
	"math/rand"
)

//生成器,在实际编程中,常见的应用场景就是调用全局的生成器,生成各项内容

//1.最简单的待缓冲生成器

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
func main() {
	randnum := GenerateIntA()

	for i := 0; i < 20; i++ {

		fmt.Println(<-randnum)
	}

}
