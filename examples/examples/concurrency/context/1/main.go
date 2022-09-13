package main

import (
	"context"
	"fmt"
	"time"
	"sync"
)


func test2(ctx context.Context, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Println("test2 is finished")
			if v := ctx.Value("one"); v != nil {
				fmt.Println("found value:", v)
			}
			return
		}
	}
}


func test(ctx context.Context, wg *sync.WaitGroup, duration time.Duration) {
	ctxInner, cancel := context.WithCancel(ctx)
	wg.Add(1)
	defer wg.Done()

	go test2(ctxInner, wg, duration * 2)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("test is finished")
			return
		case <-time.After(duration):
			cancel()
			return
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	
	// type myString string
	// one := myString("111")
	ctx = context.WithValue(ctx, "one", 1)

	wg.Add(1)
	go test(ctx, wg, time.Second * 3)
	
	<- time.After(time.Second * 1) //5 1
	cancel()

	wg.Wait()
}