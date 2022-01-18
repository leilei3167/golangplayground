package main

import (
	"fmt"
	"reflect"
)

/* 反射最佳实践：使用反射来获取结构体的字段，调用结构体的方法，并获取结构体的标签的值 */
/* 两个重要的Value方法，Methond和Call */
//定义一个结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float64
	Sex   string
}

//绑定方法显示s的值
func (s Monster) Print() {
	fmt.Println("-----strt-----")
	fmt.Println(s)
	fmt.Println("-----end-----")
}

//返回两个数的和
func (s Monster) Getsum(n1, n2 int) int {
	return n1 + n2

}

//接收4个值 给monster赋值
func (s Monster) Set(name string, age int, score float64, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex

}
func testStruct(a interface{}) {
	//获取reflect.Type类型和reflect.Value类型
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	//获取到a对应的反射类别
	kd := val.Kind()
	//判断val的Kind是否是结构体
	if kd != reflect.Struct {
		fmt.Println("错误")
		return //退出函数
	}
	//获取结构体有几个字段
	num := val.NumField()
	fmt.Printf("该结构体有%d个字段\n", num) //返回4
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("字段[%d]的值为%v\n", i, val.Field(i))
		//获取标签，注意这里需要通过reflect.Type的方法来获取
		tagVal := typ.Field(i).Tag.Get("json") //***此处不能使用val.field！！！
		if tagVal != "" {
			fmt.Printf("字段[%d]的tag为%v\n", i, tagVal)
		}

	}

	numofstruct := val.NumMethod()          //返回有多少方法
	fmt.Printf("该结构体有%d个方法\n", numofstruct) //返回3
	val.Method(1).Call(nil)                 //获取到第二个方法，并且调用。注意方法排序是按照函数名的ascii码排序的！！

	//调用结构体的第1个方法（method[0]）
	var params []reflect.Value //因为Call方法要求输入flect.Value类型的切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是[]reflect.Value,返回的也是，详见官方文档Call
	fmt.Println("res=", res[0].Int()) //因为res是个切片

}

//获取到该结构体有多少个方法

func main() {

	a := Monster{Name: "黄鼠狼", Age: 40, Score: 5.6} //创建一个monster的实例
	//将其传递给函数
	testStruct(a)
}
