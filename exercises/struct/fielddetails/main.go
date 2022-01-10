package main

import "fmt"

type Person struct { //创建一个结构体Person
	Name  string
	Age   int
	Score [5]float64
	ptr   *int              //指针
	slice []int             //切片
	map1  map[string]string //map
}

/* 结构体中，字段的初始值和其类型是相同的，引用类型都是nil
用之前必须先make
结构体是值类型的，多个结构体之间的字段相互独立，互不影响
*/

func main() {
	//定义一个结构体变量
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3") //未指定空间都是nil值
	}
	p1.slice = make([]int, 4)
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "mother"

	fmt.Println(p1)

}
