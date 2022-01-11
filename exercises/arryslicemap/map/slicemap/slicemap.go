package main

import "fmt"

func main() {
	//声明一个map类型的切片并给切片分配空间
	monsters := make([]map[string]string, 2) //准备放入两个值

	//增加信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string) //给map分配空间
		monsters[0]["name"] = "牛魔王"
		monsters[0]["年龄"] = "500"

	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string) //给map分配空间
		monsters[1]["name"] = "alahs"
		monsters[1]["年龄"] = "250"

	}//此刻已满2个，如果按这种方式再增加就会报错，应该使用切片的append再增加
	newmonster:=map[string]string{"name":"新妖怪xxx","age":"200"}
	monsters=append(monsters,newmonster)//设定一个新的map 将其通过append加入之前的map切片

	fmt.Println(monsters)
}
