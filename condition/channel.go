package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	
	go sendData(ch)
	go getData(ch)

	//这里省略time.Sleep会让其他goruntine不执行，因为来不及执行goruntine就退出了main函数
	time.Sleep(1000 * time.Millisecond)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Printf("%s ", input)
	}
}