package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/* 创建一个学生结构体的集合，并用分数排序 */
type student struct {
	Name  string
	Age   int
	Score int
}

//使用sort给结构体或切片排序要实现3个方法
func (a stuslice) Len() int { //返回长度
	return len(a)

}
func (s stuslice) Less(i, j int) bool { //比较元素大小
	return s[i].Score > s[j].Score

}

func (s stuslice) Swap(i, j int) {
	s[i].Score, s[j].Score = s[j].Score, s[i].Score //交换元素顺序

}

type stuslice []student

func main() {
	rand.Seed(time.Now().UnixNano())
	var a stuslice           //声明a是stuslice类型的切片，里面装的是Student结构体
	for i := 0; i < 5; i++ { //通过循环来创建结构体
		b := student{
			Name:  fmt.Sprintf("stu%d", rand.Intn(100)), //串联所有输出生成并返回一个字符串
			Age:   rand.Intn(20),
			Score: rand.Intn(150), //随机生成0-149的整数
		}
		a = append(a, b)
	}
	fmt.Println("排序前的数组")
	for _, v := range a {
		fmt.Println(v)
	}
	//给结构体排序要用到sort，要求实现三个方法，less，len，swap

	sort.Sort(a)
	fmt.Println("排序后的数组")
	for _, v := range a {
		fmt.Println(v)
	}

}
