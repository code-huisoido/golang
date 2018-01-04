package main

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	//如果是:=，则下面的a输出0，如果是=，则下面的a输出为10
	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为: %d\n", a)
	}

	if a < b {
		fmt.Printf("a的值为: %d\n", a)
	}

	for index,value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}
}