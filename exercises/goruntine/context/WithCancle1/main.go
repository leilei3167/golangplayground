package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/* 在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。
请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。
用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，
比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，
所有用来处理该请求的 goroutine 都应该迅速退出，
然后系统才能释放这些 goroutine 占用的资源。 */

//退出管理机制
//一.用全局变量的方式(设置一个单独的全局变量exit,需要退出时设置为false),缺点:
//1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。

//二.用channel,函数参数设置为传入一个通知chan
//使用select监听一个无缓冲的chan,当需要关闭是向其中放一个值,select马上能感知,执行
//对应的case来终止
//缺点:需要维护一个公用的chan

//三:用context优雅的实现
var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx) //协程再开协程时,只需传入ctx
Loop:
	for {

		select {
		case <-ctx.Done(): //等待上级发通知
			break Loop
		default:
			fmt.Println("worker1 working")
			time.Sleep(time.Second * 1)
		}

	}
	wg.Done()
}

func worker2(ctx context.Context) {
Loop:
	for {

		select {
		case <-ctx.Done(): //等待上级发通知
			break Loop
		default:
			fmt.Println("worker2 working")
			time.Sleep(time.Second * 1)
		}

	}
	wg.Done()
}

func worker3() {
	for {
		fmt.Println("worker3 is working")
		time.Sleep(time.Second)

	}

}

func main() {
	//创建root context
	ctx, cancle := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	wg.Add(1)
	go worker3()//此协程不收影响
	time.Sleep(5 * time.Second)
	fmt.Println("i want you all stop")
	cancle() //向此root ctx引用的协程发出退出通知
	wg.Wait()
	time.Sleep(time.Second * 5)
	fmt.Println("done")

}
