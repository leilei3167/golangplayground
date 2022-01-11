package main

import (
	"GOproject/src/go_code/exercises/struct/factory/model"//引入自己创建的model包
	"fmt"
)
/*  */
func main() {
	//创建Student的实例
	var stu = model.Student{//包名.结构体名
		Name:  "tom",
		Score: 78.9,
	}
	fmt.Println(stu)
}
