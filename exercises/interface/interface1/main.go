package main

import "fmt"

/* 在go中，接口是一组方法签名，当某个类型为这个接口中的所有方法提供了方法的实现，
他就被称为接口,接口可以减少耦合 */
//声明一个接口
type Usb interface {
	Start() //没有返回值就不写
	Stop()
	//声明了两个没有实现的方法
}
type Phone struct {
}

//让Phone实现接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

//让Camera也实现

type Camera struct {
}

func (p Camera) Start() {
	fmt.Println("相机开始工作")
}

func (p Camera) Stop() {
fmt.Println("相机停止工作") 
}

type Computer struct {
}

//编写一个方法working，只能接收一个Usb接口类型的变量（指实现了Usb接口中声明的所有方法）
//此刻相机和手机都可以是Usb接口类型的变量
func (c Computer) Working(u Usb) { //此处体现了多态
	u.Start()
	u.Stop()
	//通过接口变量调用start 和stop方法
}

func main() {
	//测试,先定义结构体
	computer := Computer{}
	//phone := Phone{}
	camera := Camera{}
	computer.Working(camera) //因为phone和camera此刻都是Usb接口类型，满足传递条件

}
