package main

import (
	"fmt"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "测试http")

}

func main() {

	http.HandleFunc("/http", handler1)

	http.ListenAndServe(":8080", nil)

}
