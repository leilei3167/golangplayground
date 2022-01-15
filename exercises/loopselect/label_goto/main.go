package main

import "fmt"

func main() {
	/* 标签一般全为大写，标签可以与选择循环语句配合，一般加在continue和goto语句之后
	   实际工作中应尽量避免使用，即使要使用也建议使用证向的，即标签在语句之后的 */
	a := 1
	b := 9
	goto TARGET
	a = 2
	b = 4

TARGET:
	b += a
	fmt.Printf("a is %v   b is %v", a, b)

}
