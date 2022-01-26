package main

import "fmt"

func main() {

	arr := []int{2, 3, 5, 12, 54, 65}
	for k, v := range arr {
		if v == 12 {
			fmt.Println("12的下标是%d\n", k)
		}
		fmt.Printf("arr[%d]=%v\n", k, v)
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {

		fmt.Printf("%s-->%s\n", k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}
