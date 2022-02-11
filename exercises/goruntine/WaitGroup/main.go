package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

//设置全局的WaitGroup

var wg sync.WaitGroup

func Geturl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	time.Sleep(time.Second * 2)
	wg.Done()
}

func main() {
	urls := []string{"https://www.baidu.com", "https://www.qq.com"}
	for _, url := range urls {
		wg.Add(1)
		go Geturl(url)

	}
	wg.Wait()

	fmt.Println("done")
}
