package main //求一个数组中的最大值，并得到对应的下标
import "fmt"

//1.声明一个数组
//2.假定第一个元素就是最大值，下标就是0
//3.然后从第二个元素开始循环比较，有最大值则交换
func main() {
	arr := [...]int{1, -9, 9, 23, 11}
	maxval := arr[0]
	maxindex := 0
	for i := 1; i < len(arr); i++ {
		if maxval < arr[i] {
			maxval = arr[i]
			maxindex = i

		}

	}
	fmt.Printf("最大值为%d，下标为%d\t", maxval, maxindex)
}
