package main

/* atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。
这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。
应通过通信来共享内存，而不通过共享内存实现通信。 */
import (
	"fmt"
	"runtime"

	"sync/atomic"
	"time"
)

//go中主要是通过通道间的沟通来实现状态管理的,也可以通过sync/atomic包来
//实现多个Gox协程的原子计数

func main() {
	var ops uint64 = 0

	//模拟并发更新,启动50个协程
	for i := 0; i < 50; i++ {
		go func() {

			for {
				// 使用 `AddUint64` 来让计数器自动增加，使用
				// `&` 语法来给出 `ops` 的内存地址。
				atomic.AddUint64(&ops, 1)

				// 允许其它 Go 协程的执行
				runtime.Gosched()

			}

		}()

	}
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops) //原子性的获取ops的值
	fmt.Println("ops:", opsFinal)

}
