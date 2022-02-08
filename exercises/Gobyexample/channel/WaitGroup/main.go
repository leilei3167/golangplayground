package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/* goroutine和channel一个用于并发一个用于通信,无缓冲通道还具有同步的功能
sync包中也提供了多个协程同步的机制,主要是通过WaitGroup实现的 */

var wg sync.WaitGroup

//WaitGroup结构体有三个方法
//func (wg *WaitGroup)Add(delta int)   添加一个等待信号
//func (wg *WaitGroup)Done()           释放一个等待信号
//func (wg *WaitGroup)Wait()            等待

var urls = []string{
	"http://www.golang.org/",
	"http://www.baidu.com/",
	"http://www.qq.com/",
}

func main() {
	for _, url := range urls {
		//遍历切片,为每一个url启动一个协程,同时给wg加1

		wg.Add(1)

		go func(url string) {
			//一个协程完成后给wg-1,wg.Done等价于wg.Add(-1)
			defer wg.Done()

			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)

			} else {
				log.Fatal(err)
			}

		}(url)

	}
	//等待所有请求结束

	wg.Wait()

}
