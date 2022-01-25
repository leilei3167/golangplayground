package main

import "fmt"

func SelectSort(arr []int) {
	n := len(arr)
	//n-1轮比较
	for i := 0; i < n-1; i++ {
		//假设最小的是arr[i]
		min := arr[i]
		minindex := i
		//从i+1开始往后选择最小值

		for j := i + 1; j < n; j++ {
			if arr[j] < min { //如果找到比min更小的，将它设置为min
				min = arr[j]
				minindex = j

			}

		}
		if i != minindex {
			arr[i], arr[minindex] = arr[minindex], arr[i]
		}
	}
	fmt.Println(arr)

}
func main() {
	list := []int{4, 2, 3, 1, 5, 9, 1}
	SelectSort(list)

}
