package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	//m sync.Mutex
	n uint64
}

var wg sync.WaitGroup

func (c *Counter) Value() uint64 {
	//c.m.Lock()
	//defer c.m.Unlock()
	return c.n
}

func (c *Counter) Increase(delta uint64) {
	//c.m.Lock()
	c.n += delta
	//c.m.Unlock()
}

func main() {
	var c Counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for k := 0; k < 100; k++ {
				c.Increase(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.Value()) // 10000
}
