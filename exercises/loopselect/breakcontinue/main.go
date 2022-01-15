package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break
			}
			fmt.Println("j=:", j)

		}
		fmt.Println("i=:", i)
	}
	//continue的使用：会跳过循环内剩余代码，执行下一次循环
	for i := 0; i < 10; i++ {
		if i == 4 || i == 6 || i == 8 {
			continue//满足条件的会跳过接下来的print 结果就是4 6 8被跳过
		}//continue可以和标签配合使用
		fmt.Println(i)

	}

}
