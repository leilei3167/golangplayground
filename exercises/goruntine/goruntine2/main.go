package main

import (
	"fmt"
	"sync"
	"time"
)

/* 计算1--20的各个数的阶乘，并把各个数的阶乘放入到map中，最后显示出来 */

//1.编写一个函数，来计算各个数的阶乘，并放入到map中
//2.我们启动的协程多个，统计的将结果放入到map中
//3.map需要做成全局的，因为不同协程都要访问

var mymap = make(map[int]int, 10)

//声明一个全局互斥锁,sync全称synchorized（同步）
var lock sync.Mutex

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	//加锁
	lock.Lock()
	mymap[n] = res //将阶乘结果放入map，没有保护机制，会导致并发写入map错误
	lock.Unlock()  //解锁
}

func main() {
	//开启200个协程
	for i := 1; i <= 20; i++ {
		go test(i)

	}
	time.Sleep(3 * time.Second)

	//输出结果
	lock.Lock()
	for k, v := range mymap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
	lock.Unlock()
}
