package main

import "fmt"

//创建Person类型

func main() {
	p1 := Person{name: "wangergou", age: 30}
	fmt.Println(p1.name, p1.age)
	p1.eat()

	s1 := Student{Person{name: "leilei", age: 15}, "成都理工大学"}
	fmt.Println(s1.name)   //s1.Person.name
	fmt.Println(s1.age)    //子类对象可以直接访问父类的字段
	fmt.Println(s1.school) //子类对象，访问自己新增的字段属性

	s1.Person.eat() //子类访问父类已有的方法
	s1.study()      //子类对象访问自己的方法
	s1.eat()        //如果存在方法重写，子类对象访问重写的方法
}

//1.定义一个“父类
type Person struct {
	name string
	age  int
}

//2.定义一个”子类
type Student struct {
	Person //结构体嵌套，模拟继承性
	school string
}

//3.方法
func (p Person) eat() {
	fmt.Println("父类的方法，吃窝窝头")
}
func (s Student) study() {
	fmt.Println("子类的方法，学习")
}
func (s Student) eat() {
	fmt.Println("子类重写的方法，吃kfc")
}
