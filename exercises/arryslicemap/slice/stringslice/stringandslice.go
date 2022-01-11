package main
import "fmt"
/* 字符串string底层是一个byte数组 因此可以进行切片处理*/
func main()  {
	str:="hello@atguigu"
	//使用切片获取某个字段
	slice :=str[6:]
	fmt.Println(slice)
	/* string是不可改变的，如果确实要修改，就必须将string转换成切片 */
	
}
