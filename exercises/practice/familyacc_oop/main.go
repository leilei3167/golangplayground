package main

import (
	"fmt"
)

func main() {
	//声明一个变量，用来接收客户选择
	key := ""
	//声明一个变量控制是否退出循环
	loop := true
	//定义变余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次的说明
	note := ""
	//收支详情用字符串记录，当有收支时，只需对他进行拼接
	detail := "类型\t账户金额\t收支金额\t说  明"
	//可以定义一个变量，记录是否出现过收支记录
	flag := false
	for {
		fmt.Println("\n---------------家庭收支记账软件----------------")
		fmt.Println("                  1.收支明细")
		fmt.Println("                  2.登记收入")
		fmt.Println("                  3.登记消费")
		fmt.Println("                  4.退出软件")
		fmt.Print("请选择<1-4>")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if !flag {
				fmt.Println("-----------------当前收支明细------------------")
				fmt.Println("当前还没有收支记录，现在就登记一笔吧")

			} else {
				fmt.Println("-----------------当前收支明细------------------")
				fmt.Println(detail)
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入  \t%v\t\t%v\t\t%v", balance, money, note)
			fmt.Printf("****记录完成！当前账户余额：%v：", balance)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			for balance < money {
				fmt.Println("超过当前余额！请重新输入金额")
				fmt.Scanln(&money)
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出  \t%v\t\t%v\t\t%v", balance, money, note)
			fmt.Printf("****记录完成！当前账户余额：%v：", balance)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？ y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，重新输入！")
			}
			if choice == "y" {
				loop = false
			}

		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你已成功退出！")

}
