package main

import (
	"fmt"
	"time"
)

/* 我们可以使用通道来同步协程间的执行状态
这个例子使用阻塞的接收方式来等待一个协程执行结束 */

//done通道将被用于通知其他协程这个函数已经工作完毕

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 3)
	fmt.Println("done")

	//向通道发送一个值表示已经完工
	done <- true

}

func main() {

	done := make(chan bool)

	go worker(done)

	//程序将在收到worker发出通知前一直阻塞
	fmt.Println("等待worker写入bool")
	<-done //如果把这一步移除,主协程立刻就会执行完毕
	fmt.Println("此刻已取出done中的bool")
}
