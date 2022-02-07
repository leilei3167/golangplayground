package main

import "fmt"

/*
goroutine在执行上来说是轻量级的线程


*/
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)

	}

}

func main() {
	//一般的调用函数,是以同步的形式调用的,同步(synchronously)
	//是所有的操作都做完，才返回给用户结果
	f("ahahahahhahaha")

	//而使用go 关键字开启协程,这个协程将会并发的执行这个函数

	go f("goroutine")

	//也可以为匿名函数启动协程

	go func(msg string) {
		fmt.Println(msg)
	}("no-namefunction")

	//以上两个goroutine在独立的协程中异步(asynchronously)运行
	fmt.Println("按回车键结束程序")
	fmt.Scanln()
	fmt.Println("done")

}
