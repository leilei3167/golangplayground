package main

import "fmt"

func test() func() int {
	i := 0
	return func() int {
		i++
		return i

	}

}

func main() {
	a := test()      //a此刻为返回的匿名函数
	fmt.Println(a()) //表示调用a这个匿名函数
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("--------再来一个b")
	b := test()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println("再来个a：", a())

}
