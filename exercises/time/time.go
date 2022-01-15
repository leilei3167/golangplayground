package main

import (
	"fmt"
	"time"
)

func main(){

//1.获取当前时间
	t1:=time.Now()
	fmt.Println(t1)

//2.获取指定时间
t2:=time.Date(2008,7,18,18,54,23,0,time.Local)
fmt.Println(t2)

//3.time->string
s1:=t1.Format("2006年1月2日 15:04:04")
fmt.Println(s1)
}