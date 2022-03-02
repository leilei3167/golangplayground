package service

import "golangplayground/exercises/practice/customerMannage_oop/model"

//该结构体完成对Customer的操作，包括增删改查
type CustomerService struct {
	//创建切片来实现客户储存，因为客户的数量是动态变化的
	customers []model.Customer
	//声明一个字段 表示当前切片含有多少个客户
	customerNum int //记录客户id，可作为新客户的id+1

}

//编写一个方法，可以返回一个CustomerService实例，这样在view里才好去调用

func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	cusomerServie := &CustomerService{}                                   //此处是对应字段中的切片，给切片分配内存，等价于customer=new（CustomerService）
	cusomerServie.customerNum = 1                                         //字段中的Num
	customer := model.Newcustomer(1, "张三", "男", 20, "112", "zs@sohu.com") //用model的方法实例化一个客户
	cusomerServie.customers = append(cusomerServie.customers, customer)   //往service结构体里的切片加客户
	return cusomerServie

}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers

}
