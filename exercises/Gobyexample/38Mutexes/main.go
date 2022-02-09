package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()         //上锁
	defer c.mu.Unlock() //解锁
	c.counters[name]++

}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncreament := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)//多个协程访问同一个Container,不加锁将会导致deadlock

		}
		wg.Done()

	}
	wg.Add(3)

	go doIncreament("a", 10000)
	go doIncreament("a", 10000)
	go doIncreament("b", 10000)
	wg.Wait()
	fmt.Println(c.counters)
}
