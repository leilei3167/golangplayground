package main

import "fmt"

func InsertSort(arr []int) {
	n := len(arr)
	if n <= 2 {
		panic("请输入2个以上的数组")
	}
	//插入排序，从第二个元素开始作为待排序数，与左边的数进行比较
	for i := 1; i < n; i++ {
		wait := arr[i]                       //待排序的数从第二个开始
		j := i - 1                           //待排数前一个数的下标
		for ; j >= 0 && arr[j] > wait; j-- { //在已排序数据中找有没有比待排序更小的

			arr[j+1] = arr[j] //左边的数如果发现比待排更大的，则左移

		}
		arr[j+1] = wait

	}
	fmt.Println(arr)
}
func main() {
	a := []int{3, 4, 1, 2}
	InsertSort(a)
}
