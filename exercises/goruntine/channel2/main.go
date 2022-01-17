package main

import "fmt"

/* 要创建一个能够放入任何类型的channel，必须借助空接口 */
type Cat struct {
	name string
	age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat
	//如果只想获得第三个元素，则必须只有先将前两个推出
	<-allChan
	<-allChan
	newcat := <-allChan
	fmt.Printf("类型：%T，值：%v\n", newcat, newcat) //类型：main.Cat，值：{小花猫 4}
	//fmt.Printf("newcat.name=%v", newcat.name) //此方法是错的，因为在编译层面还会将newcat识别为空接口类型，应先使用类型断言
	a := newcat.(Cat) //类型断言，newcat是否是Cat类型
	fmt.Printf("newcat.name=%v\n", a.name)

}
