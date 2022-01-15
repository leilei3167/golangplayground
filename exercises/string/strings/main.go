package main

import (
	"fmt"
	"reflect"
	"strings"
)

type a struct {
	Name string
}

func main() {
	var tempA = new(a) //new创建结构体 和直接创建结构体
	(*tempA).Name = "asd"

	var tempB = &a{
		Name: "asd",
	}

	fmt.Println(reflect.TypeOf(tempA))
	fmt.Println(reflect.TypeOf(tempB))
	return
	//HasPrefix 判断是否以某个字符开头，返回布尔类型
	str := "this is an example of字符串"
	fmt.Println(strings.HasPrefix(str, "th"))
	fmt.Println(strings.HasSuffix(str, "串"))
	//HasSuffix判断是否以某个字符结尾

	//Index判断指定的字符串在父字符串中的位置
	str2 := "hi i m mac"
	fmt.Printf("the position of \"mac\"is:")
	fmt.Printf("%d\n", strings.Index(str2, "mac")) //空格也计算在内
	fmt.Printf("%d\n", strings.Index(str2, "exe")) //如果返回-1表示不包含该字符串
	fmt.Printf("the last position of \"m\"is:%d\n", strings.LastIndex(str2, "m"))

	str3 := "yeyeye you are a bitch"
	//replace可以替换字符串,返回一个新的字符串
	str4 := strings.Replace(str3, "ye", "lu", 3) //表示将str3中前两个ye替换为新的lu
	fmt.Println(str4)
	fmt.Println(strings.Count(str3, "ye")) //count统计某个字符串出现的次数

	//repeat用于将某个字符串重复指定次数并返回一个新的字符串
	str5 := strings.Repeat(str3, 4)
	fmt.Println(str5)
	//Tolower 和toupper分别将字符串转换为相应的小写大写字符
	fmt.Println(strings.ToUpper(str3))

	//字符串与slice
	//feilds默认以空格来分割字符串作为返回的slice的元素
	str6 := "i love burgerking!"
	s1 := strings.Fields(str6)
	fmt.Printf("Fields将会把空白作为分隔符将字符串分割为若干块，%v\n", s1)
	for k, v := range s1 {
		fmt.Printf("s1[%d]is %s\n", k, v) //以空格为界，分元素
	}
	//你也可以使用split来指定某个内容作为分割
	s2 := strings.Split(str6, "e") //此处自定义用e来作为元素分割
	for k, v := range s2 {
		fmt.Println(k, v)
	}
}
