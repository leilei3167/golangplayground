package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //关闭管道
	//此时无法再写入
	//intChan <- 300 //报错panic: send on closed channel

	//管道关闭后读取是可以的
	n1 := <-intChan
	fmt.Println(n1)

	/* channel的遍历，用for range，无法用for循环因为取一个len就会减1
	在遍历时：
	1.如果没关闭，会出现deadlock报错
	2.如果已关闭，则会正常遍历数据，遍历完后退出

	*/

	intChan2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan2 <- i * 2 //放入十个数据到管道

	}
	//遍历，未关闭的话会报错，但是数据已经取出，在最后报错
	close(intChan2)

	for v := range intChan2 { //管道没有下标，因此for range时只会返回一个值
		fmt.Println("for range的方式遍历V=", v)
	}

	fmt.Println("------------------第二种遍历方式-------------")
	
	for {
		v, ok := <-intChan2
		if !ok {
			break
		}
		fmt.Println("v,ok的方式遍历：", v)
	}

}
