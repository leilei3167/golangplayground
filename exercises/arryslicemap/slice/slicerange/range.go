package main

import "fmt"

func main() {

	array := [4]int{4, 1, 5, 8}
	slice := array[0:3]                     //截取下标0 1 2的值
	fmt.Println(cap(slice), len(slice))     //现在指向的是array，所以容量和其相等，超过就会重新指定数组
	slice = append(slice, 1000, 2000, 3000) //超过最大容量
	fmt.Println(cap(slice), len(slice))     //现在指向的是array，所以容量和其相等，超过就会重新指定数组

	//超过容量的话会直接将原容量翻倍
	for i := 0; i < len(slice); i++ {
		fmt.Printf("for循环法:slice[%v]=%v\n", i, slice[i])

	}
	for k, v := range slice {
	fmt.Printf("forrange:slice[%v]=%v\n", k, v)
	}

}
