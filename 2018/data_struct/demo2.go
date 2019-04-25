package main
/**
 * 多变量声明
 */
var x, y int
var ( //因式分解写法一般用于全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
	//简短写法只能在函数内部使用，全局会报错non-declaration statement outside function body
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}