package main

import (
	"fmt"
	"sync"
)

func run(once *sync.Once, f func()) {
	once.Do(f)
}

func main() {
	once := &sync.Once{}
	
	f := func() {
		fmt.Println("text")
	}

	run(once, f)
	run(once, f)
}