package main

import (
	"GOproject/src/go_code/exercises/practice/familyacc_oop/utils"
	"fmt"
)

func main() {
	fmt.Println("这个是面向对象的方式！！！")

	utils.NewFamilyAccount().MainMenu()//调用NewFamilyAccount处返回的是*FamilyAccount类型
}
