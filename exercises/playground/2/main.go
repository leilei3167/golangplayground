package main

import "fmt"

type A interface {
	ping()
	pang()
}
type B interface {
	A
	sing()
}
type St struct {
	Name string
}

func (s St) ping() {
	fmt.Println("ping!!!!")
}

func (s St) pang() {
	fmt.Println("pang!!!!")
}

func (s *St) sing() {
	fmt.Println("sing!!!!")
}

func main() {
	a := new(St)
	a.Name = "bob"
	var i interface{} = a
	fmt.Printf("a的类型：%T\n", a)
	fmt.Printf("i的类型：%T\n", i)
	//判断i是否实现

	if o, ok := i.(A); ok {
		o.ping()
		o.pang()
	}

	if o, ok := i.(B); ok {
		o.sing()
	}

}
