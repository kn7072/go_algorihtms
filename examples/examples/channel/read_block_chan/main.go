package main

import (
	"fmt"
)


func main() {
	fmt.Println()

	chBuff := make(chan int, 2)
	close(chBuff)
	v := <- chBuff
	fmt.Println(v)

	chNotBuff := make(chan string)
	close(chNotBuff)
	var v2 = <- chNotBuff
	fmt.Println(v2)


}