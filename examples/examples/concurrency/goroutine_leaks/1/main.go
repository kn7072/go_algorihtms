package main

import (
	"fmt"
	"runtime"
)

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
	fmt.Println(runtime.NumGoroutine())

	// strings := make(<-chan string)
	// strings = nil
	// fmt.Printf("%T, %v\n", strings, strings)
	// for i := range strings {
	// 	fmt.Println(i)
	// }
}