package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//设置一个超时4秒
	d := time.Now().Add(time.Second * 2)//现在开始的时间2秒后
	ctx, cancle := context.WithDeadline(context.Background(), d)

	defer cancle()

	//等待
	select {
	case <-time.After(time.Second * 3):
		fmt.Println("3秒")

	case <-ctx.Done():
		fmt.Println("ctx设定的时间到,执行退出!")
		fmt.Println(ctx.Err())
	}

}
