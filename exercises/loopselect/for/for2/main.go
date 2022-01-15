package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		switch {
		case i%3 == 0:
			fmt.Printf("fizz ")
		case i%5 == 0:
			fmt.Printf("buzz ")
		default:
			fmt.Printf("%d ",i)
		}

	}
}
/* 打印1--50的数字遇到3和5的倍数打印字符 */