package main

import (
	"io"
	"log"
	"net"
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
	//处理连接,将现在的时间每隔一秒输出
	defer conn.Close() //处理完一个请求就必须关闭
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //链接断开 终止协程
		}
		time.Sleep(time.Second)

	}

}
