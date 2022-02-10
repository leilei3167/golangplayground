package main

import (
	"fmt"
	"time"
)

/* s使用select可以解决从管道取数据的阻塞问题 */

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {

		intChan <- i

	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprint("%d", i)

	}
	//传统方法遍历时，如果不关闭管道的话将会阻塞导致死锁，在实际开发中有可能不好掌握关闭管道时机
	//可以使用select的方式解决

	for {
		select {
		case v := <-intChan: //注意：如果intChan一直没有关闭也不会一直阻塞导致deadlock
			//会自动到下一个case匹配
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(200 * time.Millisecond)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据\n%s", v)
			time.Sleep(200 * time.Millisecond)
		default:
			fmt.Println("都取完了，接下来程序员就可以进行操作了\n")
			return
		}

	}

}
