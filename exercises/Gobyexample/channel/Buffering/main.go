package main

import "fmt"

/* 默认情况下通道是无缓冲的,意味着只有对应的接收<-chan准备好了接收时
才允许发送,可缓存通道允许在没有对应接收方的情况下,缓存限定数量的值 */

func main() {
	//创建一个允许缓存2个值的字符串通道
	mes := make(chan string, 2)

	//可接收2个值
	mes <- "ABC"
	mes <- "DEF"

	fmt.Println(<-mes)
	fmt.Println(<-mes)

}
