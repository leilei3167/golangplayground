package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	// Len方法返回集合中的元素个数
	Len() int
	// Less方法报告索引i的元素是否比索引j的元素小
	Less(i, j int) bool
	// Swap方法交换索引i和j的两个元素
	Swap(i, j int)
}
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//---------------------------------------------------------------------
//---------------------------------------------------------------------
//-----------------以上部分纯粹是为了调用sort.sort函数---------------------
//---------------------------------------------------------------------

/* 启动一个协程，将1--20000的数放入一个channel中，启动8个协程，从numChan中取出数，并计算1+...+n的值
并放到resChan中，最后8个协程完成工作后，再遍历resChan，显示结果要求 res[1]=1... */

//存放数
func putNum(num chan int) {
	for i := 1; i <= 5; i++ {
		num <- i //往里放20000个数
	}
	close(num) //放完之后关闭，不关闭会怎样？会报错，死锁

}

//执行求和，写入管道，并写入退出管道

func sum(num chan int, sum chan int, exit chan bool) {
	for {
		n, ok := <-num //取数
		if !ok {
			break
		}

		sumNum := 0
		for i := 1; i <= n; i++ {
			sumNum += i
		}

		sum <- sumNum //将和放入sum chan中

	}
	//取不到数据后ok会为假，跳出最外层循环并标记一个给exit
	exit <- true
	fmt.Println("有一个协程取完了数据退出了")
}

func main() {
	numChan := make(chan int, 10000)
	sumChan := make(chan int, 20000)
	exitChan := make(chan bool, 8)

	go putNum(numChan)

	for i := 0; i < 8; i++ { //启动8个协程
		go sum(numChan, sumChan, exitChan)

	}

	//主线程要来判断何时停止
	for i := 0; i < 8; i++ {
		<-exitChan

	}
	close(sumChan)
	//关闭后遍历取结果
	n := make(IntSlice, 0) //创建一个切片
	for {
		v, ok := <-sumChan
		if !ok {
			break
		}
		n = append(n, v)

	}
	//给切片n排序
	fmt.Printf("%T\n\n", n)
	sort.Sort(n)
	for k, v := range n {
		fmt.Printf("sumChan[%d]=%d\n", k+1, v)
	} //这样打印出来的数是无序的，打印格式不满足要求
	fmt.Println("主线程执行完毕")
}
