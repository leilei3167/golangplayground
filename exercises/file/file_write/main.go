package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件 并写入5句内容
	//1.打开文件
	filepath := "/home/leilei/wsl-code/GOproject/src/go_code/test2.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()
	//准备写入5句话
	str := "hello leilei\n"
	//写入时使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)

	}
	//因为writer是带缓存的，其实还未写入磁盘，因此要通过flush函数写入磁盘
	writer.Flush()

}
