package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	// "math"
	// "math/rand"
)

func processRequest(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		
		//x, ok := <- ctx.Done()
		//Printfln("FOR x %v, ok %v", x, ok)
		select {
			case <- ctx.Done():  //x, ok := 
				ch := ctx.Done()
				x, ok := <- ctx.Done()
				if (ok) {
					Printfln("ok", ok)
				} else {
					Printfln("x %v, ok %v", x, ok)
				}
				fmt.Printf("len %v, cap %v val %v type %T\n", len(ch), cap(ch), ch, ch)
				Printfln("Stopping processing - requestcancelled", ch)
				goto end
			default:
				Printfln("Processing request: %v", total)
				total++
				time.Sleep(time.Millisecond * 250)
		}
	}
	end:
	Printfln("Request processed...%v", total)
	wg.Done()
}

func main() {
	waitGroup := sync.WaitGroup {}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("%v, %T\n", ctx, ctx)
	go processRequest(ctx, &waitGroup, 10)

	time.Sleep(time.Second * 1)
	Printfln("Canceling request")
	cancel()

	waitGroup.Wait()
}
