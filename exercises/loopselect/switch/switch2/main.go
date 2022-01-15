/* switch用来判断类型 */
package main

import (
	"fmt"
)

//用法：用来判断类型type-switch
func main() {
	var x interface{} //空接口可以装任何东西，但空接口本身是什么类型呢？？

	switch i := x.(type) { //可直接在此初始化
	case nil:
		fmt.Printf("x的类型是%T\r\n", i) //得出结果，是nil类型
	case int:
		fmt.Printf("x的类型是%T\r\n", i)
	case float64:
		fmt.Printf("x的类型是%T\r\n", i)
	case func(int) float64:
		fmt.Printf("x的类型是%T\r\n", i)
	case bool, string:
		fmt.Printf("x的类型是%T\r\n", i)
	default:
		fmt.Printf("不知道", i)

	}

	//fallthroug可以在满足条件的case执行之后再执行下一个case
	k := 1
	switch k {
	case 1:
		fmt.Println("执行完case1了哦，但我加了fallthrough")
		fallthrough
	case 2:
		fmt.Println("所以我case2被执行了，我也加个fallthrough")
		fallthrough
	case 3:
		fmt.Println("case3在此")

	}
	//用法：省略表达式的switch可以当做if else if来使用
	n := 0
	switch {
	case n > 0 && n < 10:
		fmt.Print("i>0 and i<10")
	case n > 10 && n < 20:
		fmt.Print("dsadsadsad")
	default:
		fmt.Println("dddead")

	}
	
}
