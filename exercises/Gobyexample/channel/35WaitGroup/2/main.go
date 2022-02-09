package main

import (
	"fmt"
	"sync"
	"time"
)

//等待多个协程完成,可以使用WaitGroup
//每个WaitGroup必须通过指针传递给函数

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done() //此协程结束后计数器-1
	fmt.Printf("%s worker id:%d start!\n", time.Now(), id)

	time.Sleep(time.Second * 2)
	fmt.Printf("worker id:%d done!\n", id)

}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1) //每开一个协程计数器+1
		go worker(i, &wg)

	}
	fmt.Println("阻塞等待WaitGroup计数器归零")
	wg.Wait() //阻塞,等待计数器归0
	fmt.Println("结束")
}
