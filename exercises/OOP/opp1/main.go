package main

import (
	"fmt"
)

/*
面向对象编程，设定属性，行为，符合这些定义的，可以组成对象
*/
type Account struct { //创建结构体
	AccountNo string
	Pwd       string
	Balance   float64
}

//创建方法
//1.存款
func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功，你的余额为：", account.Balance)

}

//2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功，你的余额为：", account.Balance)
}

//3.查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("账号：%v的余额为%v\n\n\n", account.AccountNo, account.Balance)

}

func main() {
	//测试
	account := &Account{
		AccountNo: "gs111111",
		Pwd:       "666666",
		Balance:   100.00,
	}
	for i := 7; i != 0; {
		fmt.Printf("请选择你想要进行的操作：\n1.查询余额 2.取钱 3.存钱 0.退出")
		fmt.Scanf("%v", &i)
		if i == 0 {
			fmt.Println("感谢你的使用!!")
			break

		}
		switch i {
		case 1:
			fmt.Println("请输入密码")
			var x string
			fmt.Scanf("%v", &x)
			account.Query(x)
		case 2:
			fmt.Println("请输入密码")
			var x string
			var y float64
			fmt.Scanf("%v", &x)
			fmt.Println("请输取款金额")
			fmt.Scanf("%v", &y)
			account.WithDraw(y, x)
		case 3:
			fmt.Println("请输入密码")
			var x string
			var y float64
			fmt.Scanf("%v", &x)
			fmt.Println("请输存款金额")
			fmt.Scanf("%v", &y)
			account.Deposite(y, x)
		}
	}
}
