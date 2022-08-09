package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println()
	ch := make(chan int)
	close(ch)

	select {
	case <- time.After(time.Second * 2):
		fmt.Println("2")
	case <- time.After(time.Second * 3):
		fmt.Println("3")
	}

	select {
	case v, ok := <- ch:
		fmt.Printf("%v %v\n", v, ok)
	}


	fmt.Println("END")
}