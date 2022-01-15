package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//用ioutil.ReadFile一次读取到位，只适合文件较小的情况
	file := "/home/leilei/wsl-code/GOproject/src/go_code/test.txt"
	content, err := ioutil.ReadFile(file) //会返回一个切片
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("file contentx:%s", content)
	//因为open和close都被封装到ReadFile函数中了

}
