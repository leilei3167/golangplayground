package main

import "fmt"

func main() {
	str := "i am leilei!"
	fmt.Printf("the length of str is:%d\n", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("第%d个字符是：%c\n", i+1, str[i])

	}

	str2 := "雷磊牛逼"
	fmt.Printf("the length of str2 is:%d\n", len(str))
	for i := 0; i < len(str2); i++ {
		fmt.Printf("第%d个字符是：%c", i+1,str2[i])
	}

}
