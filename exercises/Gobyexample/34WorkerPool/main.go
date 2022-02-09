package main

import (
	"fmt"
	"time"
)

//使用协程和通道实现一个线程池

//将并发执行多个worker,worker将从jobs处获取任务,花费1秒时间处理任务,将
//结果反馈给results
func worker(id int, jobs <-chan int, results chan<- int) {
	//只能接收的jobs,以及只能发送的results
	for j := range jobs {
		fmt.Println("worker", id, "接收任务:", j)
		time.Sleep(time.Second)
		results <- j * 2
	}

}

func main() {
	start := time.Now()
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)

	}

	//发送9个jobs,然后关闭
	for j := 1; j <= 9; j++ {
		jobs <- j

	}
	close(jobs)

	//获取结果
	for a := 1; a <= 9; a++ {
		fmt.Println("接收到结果:", <-results)
	}
	fmt.Println(time.Since(start))
	//执行时间为3秒而不是9秒,可见是并发执行的
}
