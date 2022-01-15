package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//1.打开一个文件 写入十句话
	filepath := "/home/leilei/wsl-code/GOproject/src/go_code/exercises/file/file_write/ex1_write/写入几句话.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666) //以读写方式打开，并且可以追加
	if err != nil {                                                 //增加主要就是把打开模式改为Append
		fmt.Println(err)
	}
	defer file.Close() //延迟关闭file
	//先读取原来内容并打印出来
	reader := bufio.NewReader(file)
	fmt.Println("读取原内容成功：")
	for {

		str, err := reader.ReadString('\n')
		if err == io.EOF { //表示读取到文件末尾就停止循环
			break
		}
		fmt.Printf(str)
	}

	//写
	str := "这是ex3读取后修改再加的话\n"
	writer := bufio.NewWriter(file) //写入10句话
	for i := 0; i < 5; i++ {

		writer.WriteString(str) //writestring表示写入字符串
	}
	writer.Flush()
	fmt.Println("已完成显示原内容并追加")

}
