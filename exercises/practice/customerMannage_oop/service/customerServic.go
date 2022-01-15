package sevice

import "GOproject/src/go_code/exercises/practice/customerMannage_oop/model"

//该结构体完成对Customer的操作，包括增删改查
type CustomerService struct {
	//创建切片来实现客户储存，因为客户的数量是动态变化的
	customers []model.Customer
	//声明一个字段 表示当前切片含有多少个客户
	customerNum int //记录客户id，可作为新客户的id+1

}

//编写一个方法

func NewCustomerService() *CustomerService {
	cusomerServie := &CustomerService{}
	cusomerServie.customerNum = 1
	customer := model.Newcustomer(1, "张三", "男", 20, "112", "zs@sohu.com")//用model的方法实例化一个客户
	cusomerServie.customers = append(cusomerServie.customers, customer)//往service结构体里的切片加客户
	return cusomerServie

}
