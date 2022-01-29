package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("初始值", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	l := s[3:]
	fmt.Println(l)
	towd := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		towd[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			towd[i][j] = i + j

		}

	}

}
