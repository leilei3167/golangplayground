package main

import (
	"fmt"
	"time"
)

/* 要求统计1-80000的数字中，哪些是素数？运用goruntine和channel的知识解决 */

func putNum(intChan chan int) {
	for i := 1; i <= 100000; i++ {
		intChan <- i //把i往intChan里面放

	}
	close(intChan) //放完数据关闭，在关闭之前其他协程其实也可以同步读取
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) { //需要3个管道
	//分别代表从哪里取数据，结果往哪里放，结束的状态放哪里
	//使用for循环
	var flag bool
	for {

		num, ok := <-intChan //从指定的管道里取数据

		if !ok {
			break //判断intChan里面何时取完，取完就退出外层循环
		}
		flag = true //假设所有数是素数
		//判断num是否为素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //就说明num不是素数
				flag = false //将flag置为假后退出循环
				break

			}
		}
		if flag {
			//如果没被置为假，就说明是素数，要把结果放入primeChan里面
			primeChan <- num
		}
	}
	//	fmt.Printf("有一个协程因为取不到数据退出了\n")
	//因为无法判断其他协程是否也取不到数据，因此不能关闭，向exitChan写入内容，等待即可
	exitChan <- true //完成就写一个true，但是不关闭管道

}

func main() {
	intChan := make(chan int, 10000)   //创建一个channel，用来往里装整数，只要在读取，容量小一点都可以
	primeChan := make(chan int, 50000) //用来放结果
	exitChan := make(chan bool, 16)    //用于判断几个线程是否结束

	//开启一个协程，向intChan放入1--8000的数
	start := time.Now().UnixMilli() //统计开始时间
	go putNum(intChan)

	//开启16个协程，读取intChan里面读数据，并且判断是否是素数，将结果写入primeChan
	for i := 0; i < 16; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//在主线程判断什么时候所有协程执行完毕（exitChan），并且还要解决什么时候关闭管道的问题
	go func() {
		for i := 0; i < 16; i++ { //这部分放在主线程的话影响不好，于是单开一个匿名函数协程处理
			<-exitChan //没有重要性，可以直接扔掉

		}
		end := time.Now().UnixMilli()
		fmt.Println("使用协程耗时=", end-start) //相比于test里的传统方法，用协程效率提高十倍
		close(primeChan)                  //当从exitChan里取了16个结果，就可以放心关闭primeChan，然后遍历取值了
		//exitChan不关闭的原因是管道内只有16个数据，取也只取16个
	}()
	//遍历取结果
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程执行完毕，退出")
}
