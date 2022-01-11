package main

import "fmt"

/* 编写一个函数，要求可以接手一个n int，能够将斐波那契数列放到切片中 */
func main() {
	fmt.Println("请输入一个n值")
	var n int
	fmt.Scanf("%d", &n)
	t := fbn(n)
	fmt.Println(t)

}

func fbn(n int) []uint64 {
	array := make([]uint64, n)
	//第一个第二个都为1
	array[0] = 1
	array[1] = 1
	for i := 2; i < n; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array
}
