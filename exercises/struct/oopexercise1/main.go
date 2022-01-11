package main

import "fmt"

/*面向对象编程示例：
1.编写一个结构体
2.给结构体声明一个方法
3.在main中创建相应结构体变量，并访问方法，将结果打印出

*/
type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (student *Student) say() string { //声明方法，与*Student类型绑定
	info := fmt.Sprintf("name=%v,\ngender=%v,\nage=%v,\nid=%v,\nscore=%v\n",
		student.Name, student.Gender, student.Age, student.Id, student.Score)
	return info
}

/* 创建一个结构体，包含长宽高 */
type Box struct {
	l float64
	w float64
	h float64
}

func (a *Box) getvol() float64 {
	return a.l * a.w * a.h//声明一个方法，计算体积

}

func main() {
	var stu = Student{
		Name:   "mother",
		Gender: "male",
		Age:    19,
		Id:     1234141,
		Score:  89.99,
	}
	fmt.Println((&stu).say())

	var a Box
	fmt.Println("请分别输入长宽高，空格键分隔")
	fmt.Scanf("%v %v %v", &a.l, &a.w, &a.h)//格式化输入长宽高
	fmt.Println((&a).getvol())//调用方法
}
