package main

import "fmt"

func main() {
	var grade string = "8"
	var marks int
	fmt.Println("请输入你的成绩：")
	fmt.Scanln(&marks)
	switch {
	case marks >= 90:
		grade = "A"
	case marks < 90 && marks > 70:
		grade = "B"
	case marks <= 70:
		grade = "C"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀！")
	case grade == "B":
		fmt.Printf("一般！")
	case grade == "C":
		fmt.Printf("差！")
	default:
		fmt.Println("菜鸟")

	}

}
