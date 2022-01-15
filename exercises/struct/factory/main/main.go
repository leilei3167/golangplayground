package main

import (
	"GOproject/src/go_code/exercises/struct/factory/model" //引入自己创建的model包
	"fmt"
)

/*  */
func main() {
	//通过创建的工厂来创建Student的实例
	var tom = model.Newstudent("tom", 77.2) //此时tom是一个指针

	fmt.Println("name=", tom.Name, "score=", tom.Getscore())//将前面的score改为小写后将不能直接访问
	//有了那个Newstudent工厂，我可以在任意包内创建student的实例
	var bob = model.Newstudent("bob", 66.2)
	bob.Name = "bob2" //注意此处等价于(*bob).Name，go做了自动转换
	fmt.Println(*bob)
}
