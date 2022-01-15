package main

import (
	"fmt"
	"strings"
)

/* 当一个返回值为另一个函数的函数可以被称之为工厂函数
特别利于创建多个相似的函数

*/
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { //如果不是以suffix结尾的话，就在后面加suffix
			return name + suffix
		}
		return name
	}

}

/* 此刻MakeAddSuffix就是一个工厂函数，根据指定的suffix的不同，制造出增加不同内容的函数 */

func main() {
	addavi := MakeAddSuffix(".avi")//func addavi(name string)string{}
	addjpg := MakeAddSuffix(".jpg")

	fmt.Println(addavi("你好"))
	fmt.Println(addjpg("你好"))

}
