package main
/* 此例体现接口的作用，sort.sort函数调用条件就是参数实现其接口
对调用者来讲不需要知道背后原理，只需要满足其方法，即可
*/
import (
	"fmt"
	"math/rand"
	"sort"
)

//声明结构体
type Hero struct {
	Name string
	Age  int
}

//声明一个hero结构体切片
type Heroslice []Hero

//实现接口，Heroslice实现了以下三个方法，因此可以调用sort.sort函数
func (hs Heroslice) Len() int {
	return len(hs)
}
func (hs Heroslice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age //>为降序排列，反之为升序
}
func (hs Heroslice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp //在go中可简写为hs[i],hs[j]=hs[j],hs[i]
}

func main() {
	//先定义数组/切片
	var intslice = []int{0, -1, 10, 7, 90} //一般的切片或数组怎么排序？
	//排序，1冒泡排序2.或者sort函数
	sort.Ints(intslice)
	fmt.Println(intslice)
	fmt.Println("-----------排序前的切片------------------------")
	//对一个结构体切片排序
	//1.冒泡排序2.系统方法

	var heros Heroslice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄no.%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	//排序前的顺序
	for _, v := range heros {
		fmt.Println(v)
	}
	//调用sort.sort
	fmt.Println("--------排序后-----------------")
	sort.Sort(heros) //此函数使用条件就是heros切片实现其要求的三个方法，实现后才能调用，具体参考文档
	for _, v := range heros {
		fmt.Println(v)
	}
}
