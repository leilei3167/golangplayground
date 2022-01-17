package main

import (
	"fmt"
)

func writedata(intChan chan int) {
	//写入50个int
	for i := 0; i <= 50; i++ {

		intChan <- i
	}
	close(intChan) //关闭管道
}

func readdata(intChan chan int, exitChan chan bool) {
	//遍历取出所有数据
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readdata读到数据=%v", v)

	}
	//读取完数据后，即任务完成后，往exit管道写入一个数据
	exitChan <- true
	close(exitChan)//马上关闭
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writedata(intChan)
	go readdata(intChan, exitChan) //开启两个协程后，要增加一个判断退出时机的机制，否则主函数瞬间执行完毕
	//time.Sleep(5 * time.Second) //用sleep当然也能解决问题

	//用一个无限循环读取exitChan中的数据，有的话就可以退出继续执行主函数
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
