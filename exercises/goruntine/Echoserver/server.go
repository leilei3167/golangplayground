package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue //链接失败打印链接继续循环

		}
		//建立连接后,开启协程处理请求
		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn) //从链接中扫描数据到没有为止
	for input.Scan() {
		go echo(conn, input.Text(), time.Second) //接收到一段就开个协程

	}
	conn.Close()

}

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
	time.Sleep(delay)

}
