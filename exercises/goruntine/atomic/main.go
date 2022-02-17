package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//mutexAdd() //测试加锁速度24ms
	atomicAdd()  //测试不加锁而使用原子操作写入 速度 25ms
	atomicLoad() //测试原子性读取
	CAS()        //测试CAS
}

//加锁实现并发安全
func mutexAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a += 1
			mu.Unlock()

		}()

	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Milliseconds()
	fmt.Printf("use mutex a is %d,takes time %v\n", a, timeSpends)

}

//原子操作实现并发安全
/* 需要注意的是，所有原子操作方法的被操作数形参必须是指针类型，
通过指针变量可以获取被操作数在内存中的地址，从而施加特殊的CPU指令，
确保同一时间只有一个goroutine能够进行操作 */

//原子性的写入
func atomicAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()

	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Milliseconds()
	fmt.Printf("use mutex a is %d,takes time %v\n", a, timeSpends)

}

//原子性的读取
func atomicLoad() {
	var i int32 = 100
	atomic.LoadInt32(&i)
	fmt.Println(i)

}

//比较并交换(CAS)
//该操作执行前首先将确保被操作数的值未被改变,大量协程对变量进行读写操作时,将可能导致失败
//可用for循环多次尝试
func CAS() {
	var i int32 = 100
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("i:%v\n", i)

}
