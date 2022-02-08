package main

import "fmt"

//关闭通道后就不能再向这个通道发送值了

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	//启动一个协程,不断的从jobs中取数据
	go func() {
		fmt.Println("协程启动")
		for {
			j, ok := <-jobs //如果jobs关闭,则ok为FALSE

			if ok {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("从jobs接收完毕")
				done <- true
				return
			}

		}

	}()

	//向jobs中放入几个值
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("放入:", j)

	}
	close(jobs) //放完3个值之后关闭
	fmt.Println("输入jobs完成")

	<-done //无缓存通道同步

}
