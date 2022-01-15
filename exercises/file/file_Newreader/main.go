package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/home/leilei/wsl-code/GOproject/src/go_code/test.txt")
	//调用open会返回文件指针和一个err信息，因此需要两个变量来接收
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时要及时关闭文件
	defer file.Close() //为了避免忘记可以使用defer，不及时关闭会导致内存泄漏

	//1.带缓冲的形式，适用于较大文件 
	
	//创建一个*reader，他是带缓冲的,Newreader默认缓冲区为4096
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束一次
		if err == io.EOF {                  //io.eof表示文件的末尾，EOF来自于io包
			break

		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
}
