package main

import (
	"sync"
	"time"
	"fmt"
	// "math"
	// "math/rand"
	"context"
)

func processRequest(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <- ctx.Done():
			if (ctx.Err() == context.Canceled) {
				Printfln("Stopping processing - request	cancelled %v", context.Canceled)
			} else {
				Printfln("Stopping processing - deadline reached %v", context.Canceled)
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
	end:
	wg.Done()
}

func main() {
	waitGroup := sync.WaitGroup {}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	contBase := context.Background()
	ctxWC, _ := context.WithCancel(contBase)
	//ctx, _ := context.WithTimeout(context.Background(),	time.Second * 2)
	
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 1))
	time.Sleep(time.Second * 1)
	//fmt.Printf("Deadline %v %v", ctx.Deadline())
	fmt.Println(ctx.Deadline())
	fmt.Println(ctxWC.Deadline())
	go processRequest(ctx, &waitGroup, 10)

	chanCtx := ctx.Done()
	val, ok := <-chanCtx
	fmt.Println(val, ok)

	defer cancel()
	// time.Sleep(time.Second)
	// Printfln("Canceling request")
	// cancel()
	waitGroup.Wait()
}
	