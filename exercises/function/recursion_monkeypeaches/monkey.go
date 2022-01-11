package main

import "fmt"

func main() {

	fmt.Println("桃子数量为：", peach(9))

}

func peach(n int) int {
	if n>10||n<1 {
		fmt.Println("输入天数有误")
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2

	}

}
