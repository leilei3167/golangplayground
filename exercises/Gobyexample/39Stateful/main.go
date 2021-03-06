package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//在这个例子中，state 将被一个单独的 Go 协程拥有。这就
// 能够保证数据在并行读取时不会混乱。为了对 state 进行
// 读取或者写入，其他的 Go 协程将发送一条数据到拥有的 Go
// 协程中，然后接收对应的回复。结构体 `readOp` 和 `writeOp`
// 封装这些请求，并且是拥有 Go 协程响应的一个方式。

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	//用来计算执行操作的次数
	var readOps uint64
	var writeOps uint64

	//这两个通道分别将被其他协程用来发布读和写请求
	reads := make(chan *readOp) //reads是一个容纳readOp类型的通道
	writes := make(chan *writeOp)

	// 这个就是拥有 `state` 的那个 Go 协程，和前面例子中的
	// map一样，不过这里是被这个状态协程私有的。这个 Go 协程
	// 反复响应到达的请求。先响应到达的请求，然后返回一个值到
	// 响应通道 `resp` 来表示操作成功（或者是 `reads` 中请求的值）
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads: //reads取出的数据是结构体
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true

			}

		}

	}()

	// 启动 100 个 Go 协程通过 `reads` 通道发起对 state 所有者
	// Go 协程的读取请求。每个读取请求需要构造一个 `readOp`，
	// 发送它到 `reads` 通道中，并通过给定的 `resp` 通道接收
	// 结果。
	for r := 0; r < 100; r++ {
		//开启100个协程
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}

				reads <- &read //将read结构体放入reads通道
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond) //1ms

			}

		}()

	}
	// 用相同的方法启动 10 个写操作。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 让 Go 协程们跑 1s。
	time.Sleep(time.Second)
	// 最后，获取并报告 `ops` 值。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}

/* 在这个特殊的例子中，基于 Go 协程的比基于互斥锁的稍复杂。
这在某些例子中会有用，
例如，在你有其他通道包含其中或者当你 管理多个这样的互斥锁容易出错的时候。
你应该使用最自然 的方法，特别是关于程序正确性的时候。
*/
