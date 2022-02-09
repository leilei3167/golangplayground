package main

import (
	"fmt"
	"time"
)

//定时器是某一时刻执行一次使用的,而打点器则是想要在固定时间间隔重复执行准备的

func main() {
//打点器和定时器的机制有点相似：一个通道用来发送数据。
	ticker := time.NewTicker(time.Second * 1)

	go func() {
		for t := range ticker.C {
			fmt.Println("ticker.C遍历:", t)

		}

	}()

	//打点器可以和定时器一样被停止

	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("ticker stopped")

}
