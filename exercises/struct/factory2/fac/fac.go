package fac

import "fmt"

type animal struct {
	Name string
	foot int
} //创建一个anmal工厂，以便于其他包创建
func Newanimal(n string, f int) *animal {
	return &animal{
		Name: n,
		foot: f,
	}

}

type human struct {
	animal
	color string
} //创建一个工厂


func (h *human) speak() {
	fmt.Println("人可以说话！")

}
func (a *animal) move() {
	fmt.Println("动物都可以动！")
}

type v interface {
	move()
	speak()
}
