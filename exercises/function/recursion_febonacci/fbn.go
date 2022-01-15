package main

import "fmt"

func main() {
	result := 0
	var n int
	fmt.Println("请输入一个值")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		result = feb(i)
		fmt.Printf("no.%d is %d", i, result)

	}
}

func feb(n int) (res1 int) {
	if n <= 1 {
		res1 = 1

	} else {
		res1 = feb(n-1) + feb(n-2)

	}
	return
}
