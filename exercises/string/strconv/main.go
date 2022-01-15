package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "666"
	var an int
	var news string
	fmt.Printf("the size of ints is :%d\n", strconv.IntSize) //用于获取当前运行系统平台int类型所占位数
	an, _ = strconv.Atoi(str1)                               //将字符串转化为int型
	fmt.Printf("the integer is %d\n", an)
	an += 5
	news = strconv.Itoa(an) //将某个整型转化为字符串
	fmt.Printf("the news string is %s", news)
}
