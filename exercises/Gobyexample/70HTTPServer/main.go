package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "lalalala")

}

func headers(w http.ResponseWriter, r *http.Request) {
	//获取请求头信息,header是一个map,遍历取出,也可以用Get查询字段
	for k, v := range r.Header {
		fmt.Fprintln(w, k, v)

	}

}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/head", headers)
	http.ListenAndServe(":8080", nil)

}
