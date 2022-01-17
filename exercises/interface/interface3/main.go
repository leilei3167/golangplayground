package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}
type Anter interface {
	Inter
	String()
}
type St struct {
	Name string
}

func (St) Ping() {
	fmt.Println("ping")
}
func (*St) Pang() {
	fmt.Println("pang")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st
	//conmma ok法判断i绑定的实例是否实现了接口类型Inter
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}
	//判断i绑定的实例是否实现了接口Anter
	if p, ok := i.(Anter); ok {
		p.String()

	}
	//直接赋值法判断i是否就是*St类型
	s := i.(*St)
	fmt.Printf("%s", s.Name)

}
