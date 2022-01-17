package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMilli()
	for num := 1; num <= 80000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}

		}
		if flag {
			//省略打印
		}
	}
	end := time.Now().UnixMilli()
	fmt.Println("传统方法耗时=", end-start)
}
