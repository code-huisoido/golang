package main

import "fmt"

func main() {
	var grade string = "B"
	var marks int = 90

	//第一种写法:
	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default: grade = "D"
	}

	//第二种写法：
	switch {
		case grade == "A":
			fmt.Println("优秀")
		case grade == "B", grade == "C":
			fmt.Println("良好")
		case grade == "D":
			fmt.Println("及格")
		case grade == "F":
			fmt.Println("不及格")
		default:
			fmt.Println("再努力看看？")
	}

	fmt.Printf("你的等级是：%s\n", grade)
}