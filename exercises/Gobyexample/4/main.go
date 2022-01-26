package main

import "fmt"

func fun1(a int) {
	a = 100000
	return

}
func fun2(a *int) {
	*a = 200000
}
func main() {
	a := 1
	fun1(a)
	fmt.Println(a)

	fun2(&a)
	fmt.Println(a)
	
}
