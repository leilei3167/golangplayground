package main

import (
	"fmt"
	"sort"
)

func main() {
	//1.排序字符串,他会改变原来的序列而不是给定一个新值
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("排序后:", strs)

	//2.int排序

	ints := []int{4, 12, 4, 6, 1}
	sort.Ints(ints)//Ints函数将a排序为递增顺序。
	fmt.Println(ints)

	//3.检查一个序列是否已经排序
	fmt.Println(sort.IntsAreSorted(ints))

}
