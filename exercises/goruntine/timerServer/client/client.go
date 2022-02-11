package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	go fun1()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//从连接中读取数据,并输出终端 os.stdout

	_, err = io.Copy(os.Stdout, conn) //从网络链接中读取,直到到达EOF或出错
	//copy内部是一个循环
	if err != nil {
		log.Fatal(err)
	}

}

func fun1() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//从连接中读取数据,并输出终端 os.stdout
	fmt.Println("新时钟:")
	_, err = io.Copy(os.Stdout, conn) //从网络链接中读取,直到到达EOF或出错
	//copy内部是一个循环
	if err != nil {
		log.Fatal(err)
	}
}
