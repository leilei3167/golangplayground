package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest(b interface{}) { //要想对任意一个类型进行反射，就需要空接口
	//通过反射获取到传入的变量的type和kind（类别），值
	//先获取reflect.type
	rTyp := reflect.TypeOf(b)                  //返回b的类型,并可使用Type接口内的方法（参见官方文档）
	fmt.Printf("rTyp值：%v，类型：%T\n", rTyp, rTyp) //值是int，类型是*reflect.rtype

	//获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)                 //打印出100，要注意理解此处的100不是普通的100，他是一个reflect.value类型
	fmt.Printf("rVal值：%v，类型：%T\n", rVal, rVal) //打印出值是100，但要注意他是一个reflect.value类型
	//正因为typeOf返回value类型，因此返回后，其所有的针对value类型方法都可以使用（参考标准库）
	fmt.Println(rVal.Int()) //调用针对value类型的方法 取出其中的值

	//又将rVal类型转成interface{}
	iV := rVal.Interface()
	//将Interface通过断言重新转成所需的类型
	num2 := iV.(int)
	fmt.Println(num2)

}

func main() {
	//定义一个
	var num int = 100

	reflectTest(num) //把变量传递给空接口

}
