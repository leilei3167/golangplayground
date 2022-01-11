package main

import "fmt"

func main() {
	a := fbn(8)
	fmt.Println(a)
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)

	}

}
