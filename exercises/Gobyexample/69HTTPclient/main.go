package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	//`http.Get` 是创建 `http.Client` 对象并调用其 `Get` 方法的快捷方式。
	// 它使用了 `http.DefaultClient` 对象及其默认设置。
	defer res.Body.Close()        //body是readercloser接口
	fmt.Println(res.Status) //打印res中的status

	//打印body的内容

	scanner := bufio.NewScanner(res.Body)

	for scanner.Scan() { //读取到Scan为false
		fmt.Println(scanner.Text())

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
