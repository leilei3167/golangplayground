package main

import (
	"fmt"
	"reflect"
)

func reflect1(b interface{}) {
	rVal := reflect.ValueOf(b)
	rTyp := reflect.TypeOf(b) //获取Type
	fmt.Println(rTyp)
	fmt.Println(rVal.Kind()) //获取kind

	iv := rVal.Interface() //转换为interface类型

	flo := iv.(float64) //通过类型断言转换为float
	fmt.Printf("%T\n", flo)

}

func main() {
	var v float64 = 1.2

	reflect1(v)

}
