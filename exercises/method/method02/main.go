package main

import "fmt"

type circle struct { //创建结构体
	redius float64
}

func (c circle) area() float64 { //创建方法，只有circle类型的可以调用
	return 3.14 * c.redius * c.redius
}

//为提高效率，通常我们方法和结构体的指针类型相绑定，指针效率更高，因为只用传递地址

func (acd *circle) area2() float64 { //编译器会识别传递的是结构体指针类型，可以做相应的优化
	fmt.Printf("方法中b的地址是%p", acd)
	(*acd).redius = 11.11 //改变外面b的值
	return 3.14 * acd.redius * (*acd).redius
	//如此处（*acd）.redius等价于b.redius，是一个语法糖
}

//go中，所有自定义类型都可以指定方法如：
type motherfucker int //自定义一个int类型，名为motherfucker

func (i *motherfucker) changeval() {
	*i = *i + 5 //给这个自定义类型指定方法，只有motherfucker指针类型才能调用
}


func main() {
	c := circle{redius: 6.66}

	res := c.area()
	fmt.Println("面积是", res)

	b := circle{redius: 5.0}
	res2 := (&b).area2()
	//编译器做了优化，(&b).area2()也等价于b.area2()，语法糖
	fmt.Println("尝试与结构体指针绑定 值为", res2)
	fmt.Printf("主函数b的地址是：%p", &b) //可以发现和方法中的地址是相同的，因此方法中如果将b修改，那么将会使main中的值改变
	fmt.Println(b)                //输出11，证实

	var father motherfucker = 10 //定义一个叫father的变量，类型为motherfucker，值为10
	(&father).changeval()        //调用方法，可以省略取址符
	fmt.Println(father)

}
