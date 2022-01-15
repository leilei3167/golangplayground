package main

import (
	"fmt"
	"io/ioutil"
) //将一个文件的内容 加到另一个文件中
func main() {
	//1.首先将文件1读取到内存

	//2.将读取到的内容写入另一个文件
	file1 := "/home/leilei/wsl-code/GOproject/src/go_code/exercises/file/file_write/ex4_move/file1.txt"
	file2 := "/home/leilei/wsl-code/GOproject/src/go_code/exercises/file/file_write/ex4_move/file2.txt"

	content1, err := ioutil.ReadFile(file1) //outil里面的read是直接读取不带缓存，返回的是一个byte切片和一个err，需要接收
	if err != nil {
		fmt.Println("读取文件1出错！", err)

	}
	fmt.Printf("文件1的内容是：%s\n", content1)
	err = ioutil.WriteFile(file2, content1, 0666) //写入
	if err != nil {
		fmt.Println("写入文件出错！", err)

	}
	fmt.Printf("文件2追加后完成")

}
