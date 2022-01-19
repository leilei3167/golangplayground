package main

import "fmt"

/* 统计1--1000 哪些是素数？用协程和管道解决 */
var putchan = make(chan int, 1000)
var exitchan = make(chan int, 8)
var reschan = make(chan int, 1000)

func put() {
	for i := 1; i <= 1000; i++ {
		putchan <- i

	}
	close(putchan)
}

func read(put chan int, res chan int, exit chan int) {
	var flag bool
	for { //遍历取数
		o, ok := <-put
		if !ok {
			break
		}
		flag = true
		for i := 2; i < o; i++ {
			if o%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			res <- o
		}

	}
	fmt.Println("有一个协程取不到退出了")
	exit <- 1
}

func main() {
	go put()

	for i := 0; i < 8; i++ {
		go read(putchan, reschan, exitchan)

	}

	//判断什么时候exitchan8个数被取完，就说明所有协程执行完毕
	for i := 0; i < 8; i++ {
		<-exitchan

	}
	close(reschan)

	for v := range reschan {
		fmt.Println(v)
	}

}
