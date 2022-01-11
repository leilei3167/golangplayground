package main

import "fmt"
/*二分查找前提是数组要先排序
先找到中间值的下标，然后让中间下标的值和要找的数比较，如果中中间值大就说明在左边，反之在右边，都不是就说明是中间值
递归执行，找不到数值时就需要退出递归，条件为左值大于右值

*/
func Binaryfind(arr *[6]int, leftindex int, rightindex int, findval int) {
	middle := (leftindex + rightindex) / 2
	if leftindex > rightindex {
		fmt.Println("找不到")
		return
	}

	if (*arr)[middle] > findval { //成立就说明要查的值在中间标左边
		Binaryfind(arr, leftindex, middle-1, findval)
	} else if (*arr)[middle] < findval {
		Binaryfind(arr, middle+1, rightindex, findval)

	} else {
		fmt.Printf("找到了 下标为%v\n", middle)
	}

}

func main() {
	array := [6]int{1, 9, 10, 89, 1234}
	Binaryfind(&array, 0, len(array)-1, 1234)
}
