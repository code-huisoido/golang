package main

import "fmt"

/**
 * 这里要识别出k = 3<<iota，l = 3<<iota
 * 所以k = 3<<2, l = 3<<3，也就是k=1100=12，l=11000=24s
 */
const (
	i = 1<<iota
	j = 3<<iota
	k 
	l
)

func main() {
	fmt.Println("i =", i)
	fmt.Println("j =", j)
	fmt.Println("k =", k)
	fmt.Println("l =", l)
}