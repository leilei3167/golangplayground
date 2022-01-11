package main

//随机生成五个数，并将其反转打印
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [5]int
	rand.Seed(time.Now().UnixNano()) //设置当前的时间为Seed，则可以实现每次运行生成不同的随机数
	//unix单位默认是秒，如果一秒内运行两次值还会是一样的，建议用纳秒

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) //生成大于等于0小于100的随机数,因为没有设定seed每次数都会一样

	}
	fmt.Println("生成的随机数组是", arr) //打印随机生成的数组

	cout := 0
	for i := 0; i < len(arr)/2; i++ {//交换
		cout = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = arr[i]
		arr[i] = cout
	}
	fmt.Println("反转后的数组是", arr) //打印随机生成的数组

}
