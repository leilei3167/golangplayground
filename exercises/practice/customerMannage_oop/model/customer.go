package model

import "fmt"

/*
客户数据层
用来指定客户的信息字段
客户的字段推断来自于界面
*/
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

/* 因为本层内容主要是在service层进行使用，因此我们应该提供一个工厂方法，返回实例 */
func Newcustomer(id int,
	name string,
	gender string,
	age int,
	phone string,
	email string) Customer { //创建工厂模式，返回一个Customer示例
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}

}

//返回客户信息，格式化的(封装到这里，而不是在调用处挨个取出打印)

func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
