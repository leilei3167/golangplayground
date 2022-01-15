package main

import "fmt"

/* 主菜单：
当用户运行程序时，可以看到主菜单，当输入5时可以退出


*/
type customerView struct { //本文件用不用大写
	//定义必要的字段
	key  string //接收客户的输入
	loop bool   //表示是否循环显示菜单
	//增加一个字段调用custormerservice
	//customerService *service.CustomerService
}

func (this *customerView) mainMenu() { //定义一个方法，让customerView来调用
	for {
		fmt.Println("-------------客户信息管理软件-------------")
		fmt.Println("               1.添加客户")
		fmt.Println("               2.修改客户")
		fmt.Println("               3.删除客户")
		fmt.Println("               4.客户列表")
		fmt.Println("               5.退出")
		fmt.Printf("请选择1-5：") //选项设置好之后就要考虑用变量来接收输入
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			//添加客户，此处应该就是调用service文件中的某个方法来执行，service.NewcustomerService

		case "2":
			//添加客户
		case "3":
			//添加客户
		case "4":
			//添加客户
		case "5":
			//退出
			this.loop = false
		default:
			fmt.Println("输入有误 请重新输入")
		}
		if !this.loop {
			break //跳出for循环
		}

	}
	fmt.Println("你已成功退出系统！")

}

func main() {
	//主函数中创建结构体示例 并运行主菜单，创建了实例才能去调用他的方法
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.mainMenu() //调用方法的方式：实例.方法名，customer此刻就是”this“

}
