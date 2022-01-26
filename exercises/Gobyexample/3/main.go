package main

import "fmt"

func sum(num ...int) {
	fmt.Print(num, "")
	total := 0
	for _, v := range num {
		total += v

	}
	fmt.Println(total)

}

func main() {
sum(1,5,12)
}
