package main

import (
	"sync"
	"time"
	"math"
	"math/rand"
)

var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}
// var readyCond = sync.NewCond(rwmutex.RLocker())
var readyCond = sync.NewCond(&rwmutex)


var squares = map[int]int{}

func generateSquares(max int) {
	rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting condition")
	readyCond.Broadcast()
	waitGroup.Done()
}

func readSquares(id, max, iterations int) {
	readyCond.L.Lock()
	for len(squares) == 0 {
		readyCond.Wait()
	}
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key,
		squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()
	waitGroup.Done()
}
	
func main() {
	rand.Seed(time.Now().UnixNano())
	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}
	waitGroup.Add(1)
	go generateSquares(10)
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}	