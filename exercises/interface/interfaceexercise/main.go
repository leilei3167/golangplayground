package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//要求，将学生结构体切片按照分数由高到低排序
//先定义学生结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

//这些结构体的集合，需要一个切片来装
type stuslice []Student //Student是切片的类型

//要使用sort.sort函数，必须让其实现三个方法
func (s stuslice) Len() int {
	return len(s)
}
func (s stuslice) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}
func (s stuslice) Swap(i, j int) {
	s[i].Score, s[j].Score = s[j].Score, s[i].Score
}

func main() {
	var a stuslice           //声明a是stuslice类型的切片，里面装的是Student结构体
	for i := 0; i < 8; i++ { //通过循环来创建结构体
		b := Student{
			Name:  fmt.Sprintf("stu%d", rand.Intn(100)), //串联所有输出生成并返回一个字符串
			Age:   rand.Intn(20),
			Score: float64(rand.Intn(150)), //随机生成0-149的整数
		}
		a = append(a, b)
	}
	fmt.Println("------排序前的数组------")
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("------排序后的数组------")

	sort.Sort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}
