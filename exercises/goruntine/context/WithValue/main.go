package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")

	traccode, ok := ctx.Value(key).(string)
	if !ok {

		fmt.Println("invalid code")

	}
Loop:
	for {
		fmt.Printf("worker code:%s\n", traccode)
		time.Sleep(time.Second)
		select {

		case <-ctx.Done():
			break Loop
		default:

		}

	}
	fmt.Println("ctx时间到!worker done ")
	wg.Done()
}

func main() {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*2)
	//创建的timeout context中加入数值,返回一个新的context
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "1111222233333")
	//将带数据的context传入协程
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancle()
	wg.Wait()
	fmt.Println("done")
}
