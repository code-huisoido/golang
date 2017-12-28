package main

import "fmt"

func main() {
	/**
	 * 声明变量的形式
	 */
	//第一种，指定变量类型和值，也可先声明后赋值
	var a bool = true
	fmt.Println("a的值：", a)

	//第二种，根据值自动判断类型
	var b = 10
	fmt.Println("b的值：", b)

	//第三种，省略var
	c := "hello, my name is c"
	fmt.Println("c的值：", c)

	//变量默认值,bool是false，int是0，string是空
	var default_bool bool
	var default_int int
	var default_string string
	fmt.Println(default_bool, default_int, default_string)
}