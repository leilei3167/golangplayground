package main

import "fmt"

func main() {
	str := "chinese 你好 北京"
	fmt.Printf("str的长度为%d\n", len(str))
	for k, v := range str {
		fmt.Printf("第%d个字符是：%c\n", k+1, v)
	}
}
