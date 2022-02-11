package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)

	}
	defer conn.Close()

	//输出服务端返回的数据到终端
	go mustCopy(os.Stdout, conn) //从服务器接收数据
	mustCopy(conn, os.Stdin)     //终端输入数据传输到服务器

}

//服务器和客户端的信息交换
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

}
