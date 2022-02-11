package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Client struct {
	C chan string //发送数据

	Name string
	Addr string
}

//全局map,以便于存放用户
var onlineMap map[string]Client

//传递信息的管道
var message = make(chan string)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)

	}
	defer listener.Close()

	go Manager() //监听message通道,将信息取出后转发给所有的人

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handlerConn(conn)

	}

}

func handlerConn(conn net.Conn) {
	//加入新上线的用户进map

	//获取客户的网络地址
	cliAddr := conn.RemoteAddr().String()
	//链接成功一个 就为这个客户创建结构体
	cli := Client{make(chan string), cliAddr, cliAddr}

	//以地址为key,结构体为value加入map
	onlineMap[cliAddr] = cli

	//开协程,给当前客户端发送信息
	go WriteMsgToClient(cli, conn)
	//广播某人上线信息
	message <- MakeMsg(cli, "login") //会把上线消息发送至通道被manager接收
	//manager收到后会遍历map中所有在线的用户,将收到的message转发
	//提示下我是谁

	cli.C <- MakeMsg(cli, "im here")

	//新开一个协程,同时接收来自客户端的消息,写入message通道让manager转发
	//给每一个用户的通道

	//退出控制
	isQuit := make(chan bool)
	isAlive := make(chan bool)

	go func() {
		buf := make([]byte, 2048)
		for {
			n, _ := conn.Read(buf)
			if n == 0 { //说明读取0个字节 说明断开
				isQuit <- true
				// fmt.Println(err)
				return
			}

			//需要转发读到的内容
			msg := string(buf[0 : n-1])
			if len(msg) == 3 && string(msg) == "who" {
				//如果输入命令who来查询在线用户的话,遍历给当前在线列表
				conn.Write([]byte("当前在线列表:\n"))
				for _, tmp := range onlineMap {
					msg = ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli //改名后更新map
				conn.Write([]byte("rename done"))
			} else {
				message <- MakeMsg(cli, msg)
			}
			isAlive <- true //有数据交流就说明还在
		}

	}()
	//监听退出
	for {
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)              //删除当前用户
			message <- MakeMsg(cli, cli.Name+"已退出") //广播下线信息

			return
		case <-isAlive: //有内容就不会段

		case <-time.After(30 * time.Second):
			delete(onlineMap, cliAddr) //删除当前用户
			message <- MakeMsg(cli, cli.Name+"超时被下线")
			return
		}

	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {

	for msg := range cli.C { //将被manager写入的通道遍历取出并通过conn写入

		conn.Write([]byte(msg + "\n"))

	}

}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + msg

	return buf

}

//只要有消息来了,遍历map,给每个成员都发送信息
func Manager() {
	onlineMap = make(map[string]Client)

	for { //没有消息时这里会阻塞
		msg := <-message
		for _, cli := range onlineMap {
			cli.C <- msg //将收到的message写入每个结构体的通道

		}

	}
}
