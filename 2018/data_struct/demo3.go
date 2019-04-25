package main

import "fmt"
import "unsafe"

const (
	d = "abcd"
	e = len(d)
	f = unsafe.Sizeof(d)
)

func main() {
	const LENGTH int = 10
	const WIDTH = 5

	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
	println(d, e, f)
}