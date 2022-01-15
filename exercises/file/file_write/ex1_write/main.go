package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
1.openfile来自os包，以指定的选项打开文件，操作成功打开的文件就可以用于I/O
2.ioutil和bufio主要就是读写操作，前者直接读写，后者创建缓存读写，写文件要创建writer，读要创建reader
bufio要注意flush来从缓存写入

*/
func main() {
	//1.打开一个文件 写入十句话
	filepath := "/home/leilei/wsl-code/GOproject/src/go_code/exercises/file/file_write/ex1/写入几句话.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666) //打开文件，返回的*file由file接收
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //延迟关闭file
	str := "雷磊测试ex1\n"
	writer := bufio.NewWriter(file) //写入10句话
	for i := 0; i < 10; i++ {

		writer.WriteString(str) //writestring表示写入字符串
	}
	writer.Flush()
	fmt.Println("写入完成")

}
