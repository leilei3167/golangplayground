package main

import (
	"fmt"
	"reflect"
)

/* 演示反射对基本数据类型，interface{}，reflect.Value的操作 */
type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Printf("值：%v，类型：%T\n", rTyp, rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Printf("值：%v，类型：%T\n", rVal, rVal)
	//尝试获取Kind  1)rVal.Kind()   2)rTyp.Kind()  因为Type结构体和Value的方法都有可以Kind的办法
	fmt.Printf("第一种方式得出的类别：%v，第二种方式的类别：%v\n", rVal.Kind(), rTyp.Kind()) //两种方式结果一致，都为struct

	//将rVal转成interface
	iv := rVal.Interface()
	fmt.Printf("值：%v，类型：%T\n", iv, iv) //结果：值：{Bob 13}，类型：main.Student
	//此刻不能通过iv.Name来修改值
	//必须要将iv通过断言转换成想要的类型
	stu, ok := iv.(Student)
	if ok {
		stu.Name = "jack"
		fmt.Println(stu)
	}

}

func main() {
	/* //1.定义一个int
	num := 100 */
	//2.实例化一个Student

	stu1 := Student{Name: "Bob", Age: 13}

	reflectTest02(stu1)
}
