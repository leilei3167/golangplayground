package utils

import "fmt"

type FamilyAccount struct { //创建结构体 并声明必要的字段
	//声明一个变量，用来接收客户选择
	key string
	//声明一个变量控制是否退出循环
	loop bool
	//余额
	balance float64
	//每次收支的金额
	money float64
	//每次的说明
	note string
	//收支详情用字符串记录，当有收支时，只需对他进行拼接
	detail string
	//可以定义一个变量，记录是否出现过收支记录
	flag bool //判断是否是第一次使用
}

//编写一个工厂模式的构造方法，返回一个FamilyAccoun的实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		detail:  "类型\t账户金额\t收支金额\t说  明",
	}
}

//给该结构体绑定方法
//将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	if !this.flag {
		fmt.Println("-----------------当前收支明细------------------")
		fmt.Println("当前还没有收支记录，现在就登记一笔吧")

	} else {
		fmt.Println("-----------------当前收支明细------------------")
		fmt.Println(this.detail)
	}

}

//将登记收入写成方法
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收入  \t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	fmt.Printf("****记录完成！当前账户余额：%v：", this.balance)
	this.flag = true

}

//将登记支出写成方法
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	for this.balance < this.money {
		fmt.Println("超过当前余额！请重新输入金额")
		fmt.Scanln(&this.money)
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n支出  \t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	fmt.Printf("****记录完成！当前账户余额：%v：", this.balance)
	this.flag = true

}

//退出方法
func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n---------------家庭收支记账软件----------------")
		fmt.Println("                  1.收支明细")
		fmt.Println("                  2.登记收入")
		fmt.Println("                  3.登记消费")
		fmt.Println("                  4.退出软件")
		fmt.Print("请选择<1-4>")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()

		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你已成功退出！")

}
