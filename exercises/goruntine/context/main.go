package main

import (
	"context"
	"fmt"
	"time"
)

type otherContext struct {
	context.Context
}

func main() {
	//构造一个cancle context
	cxtA, cancle := context.WithCancel(context.Background())
	go worker(cxtA, "bob")

	//用withtimeout封装cxta,赋值给cxtb,这里忽略cancle
	cxtb, _ := context.WithTimeout(cxtA, time.Second*3)

	go worker(cxtb, "json") //此协程将在2秒后退出

	//使用withvalue包装cxtb
	cxtc := context.WithValue(cxtb, "aluoha", "nima")
	go workwithValue(cxtc, "mayun")

	//故意让其超时退出
	time.Sleep(time.Second * 7)
	cancle() //结束cxtA的

	time.Sleep(time.Second * 1)
	fmt.Println("main done")
}

//工作函数
func worker(ctx context.Context, name string) {
	for {
		select {
		//监听退出信息
		case <-ctx.Done():
			fmt.Printf("%s 收到退出消息!!!\n", name)
			return //终止函数
		default:
			fmt.Println(name, "正在工作!")
			time.Sleep(time.Second)

		}

	}

}

//获取通过context传递的值
func workwithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s 收到退出消息!!!\n", name)
			return
		default:
			value, ok := ctx.Value("aluoha").(string)
			if !ok {
				fmt.Println("输入的key有问题")

			}
			fmt.Println("取出值:", value)
			time.Sleep(time.Second)
		}

	}

}
