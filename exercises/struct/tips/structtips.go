package main

import (
	"encoding/json"
	"fmt"
)

/*
	1.结构体中的所有字段是连续的
	2.结构体是用户单独定义的类型，与其他类型进行转换时需要有完全相同的字段
	3.结构体用type重新定义（相当于取别名），会被认为是新数据类型，但是可以强制转换
	4.结构体每个字段上可以写一个tag，该tag可以用反射机制获取
*/
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json错误")

	}
	fmt.Println("jsonStr", string(jsonStr))

}
