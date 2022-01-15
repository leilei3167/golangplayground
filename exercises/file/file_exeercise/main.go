package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计不同类型字符的个数
type Charcount struct {
	Chcount    int
	Numcount   int
	Spacecount int
	Othercount int
}

func main() {
	//思路：打开一个文件，创建一个reader
	//每读取一行，就去统计该行有多少个英文，数字空格和其他字符，然后将结果保存到一个结构体
	file, err := os.Open("/home/leilei/wsl-code/GOproject/src/go_code/exercises/file/file_exeercise/abc.txt")
	if err != nil {
		fmt.Println("打开文件出错！", err)
	}
	defer file.Close()
	//定义一个结构体示例
	var count Charcount

	reader := bufio.NewReader(file)
	//开始循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历str统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				count.Chcount++
				fallthrough
			case v >= 'A' && v <= 'z':
				count.Chcount++
			case v == ' ' || v == '\t':
				count.Spacecount++
			case v >= '0' && v <= '9':
				count.Numcount++
			default:
				count.Othercount++

			}
		}
	}
	//输出统计结果
	fmt.Printf("字符个数为%v，数字个数为%v，空格个数为%v，其他字符%v",
		count.Chcount, count.Numcount, count.Spacecount, count.Othercount)
}
