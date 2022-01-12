package main

/*
 接口可以定义一组方法（不一定全部都被实现），接口中不能有变量，具体方法可以到某个自定义类型要使用的时候再写
 1.基本语法：
 type 接口名 interface{
	 method1(参数列表)返回值列表
	 ...
 }
接口内的方法都没有方法体（即操作部分）；接口体现了程序设计的多态和高内聚低耦合的思想；
go中的接口不需要显式的实现，只要一个变量含有接口中的所有方法，那么这个变量就实现这个接口
注意点：
1.一个自定义类型必须实现了某个接口中的所有方法，才能被称之为实现了这个接口
2.一个类型可以实现多个接口，一个接口也可以被多个类型实现
3.一个接口A可以继承多个别的接口（如B C），如果要实现A接口，那么必须将BC中的方法也全部实现
*/
/* type Binterface interface{
	test01()
}

type Cinterface interface{
	test02()
}
type Ainterface interface{//A继承了bc的方法，此刻如果要实现A，则BC中的方法全部也要实现
	Binterface
	Cinterface//但是BC中不能有重复的方法 否则会报错
	test03()
}
4.interface是一个引用类型，未初始化就是nil值
5.不含任何方法的就是空接口，任何类型都实现了空接口，即我们可以把任何一个变量赋给一个空接口类型
6.注意鉴别实现方法的类型是否是指针类型，例：
type Usb interface{
	Say()
}
type Stu struct{
}
func (s *Stu)Say(){

}
func main(){
	s:=Stu{}
	var u Usb=s//此处报错，因为实现Usb的不是Stu结构体类型，而是*Stu结构体指针类型，s前加&可解决
}
*/
