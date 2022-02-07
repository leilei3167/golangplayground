package main

import (
	"errors"
	"fmt"
)

//按照惯例,错误通常作为最后一个返回值,并且是`error`类型,是一个接口类型(只包含一个Error方法)
func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("cant work with 42")
	}

	return arg + 3, nil

}

//可以通过实现Error方法来自定义 error的类型
//这里使用自定义错误类型来表示上面例子的参数错误

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string { //只要满足这个方法,那么他就实现了error接口,可以作为error类型返回
	return fmt.Sprintf("%d-%s\n", e.arg, e.prob)

}

func f2(arg int) (int, error) {
	if arg == 42 {
		//返回自定义错误类型的值
		return -1, &argError{arg: arg, prob: "cant work with 42"}

	}

	return arg + 3, nil

}

//测试
func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil { //用切片的值调用f1
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}

	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println(r)
		}

	}

	//如果想要在程序中使用一个自定义的错误类型中的数据,
	//则需要通过类型断言,取回其类型的实例
	_, e := f2(42) //返回的是一个error类型的interface
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)

	}

}
