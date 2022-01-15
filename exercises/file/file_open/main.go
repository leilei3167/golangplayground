package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//
	file, err := os.Open("/home/leilei/wsl-code/GOproject/src/go_code/test.txt")
	//调用open会返回文件指针和一个err信息，因此需要两个变量来接收
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//输出下文件
	fmt.Printf("file=%v", file) //可以看出，返回的是个指针，file就是个指针
	//open之后一定要记得关闭
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
	PrintContent(file)
}

func PrintContent(io io.ReadWriter) {
	var buf = make([]byte, 2048)
	_, err := io.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

