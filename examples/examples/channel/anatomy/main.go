package main

import (
	"fmt"
	"runtime"
	"time"
	"context"
	"math/rand"
	"sync"
)

// https://kovardin.ru/articles/go/chanels/

func goRoutineA(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
    defer wg.Done()
	
	for {
		select {
		case <- ctx.Done():
			return
		case val := <-ch:
			fmt.Println("goRoutineA received the data", val)
		}
	}
}

func goRoutineB(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
    defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			return
		case val := <-ch:
			fmt.Println("goRoutineB received the data", val)
		}
	}
}


func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	
	wg := &sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int)
    go goRoutineA(ctx, wg, ch)
    go goRoutineB(ctx, wg, ch)
	fmt.Println(runtime.NumGoroutine())

    for i := 0; i < 10; i++ {
		ch <- rand.Intn(10)
	}
	

    <- time.After(time.Second * 5)
	cancel()

	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
}