package main

import "fmt"

func main() {
	//const常量可以仅赋予第一个常量，后续默认与第一个常量的值一样。
	const (
		var_1 = 10
		var_2
		var_3
	)
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Printf("var_1:%d, var_2:%d, var_3:%d", var_1, var_2, var_3)
	println()
	fmt.Println(a, b, c, d, e, f, g, h, i)
}