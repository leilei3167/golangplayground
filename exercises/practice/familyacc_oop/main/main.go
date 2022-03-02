package main

import (
	"golangplayground/exercises/practice/familyacc_oop/utils"
	"fmt"
)

func main() {
	fmt.Println("这个是面向对象的方式！！！")
	temp := utils.NewFamilyAccount()
	temp.MainMenu() //调用NewFamilyAccount处返回的是*FamilyAccount类型
}
