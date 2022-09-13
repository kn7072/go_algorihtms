package main

import "sync"


func doSum(count int, val *int, waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 0; i < count; i++ {
		mutex.Lock()
		*val++
		mutex.Unlock()
	}
	waitGroup.Done()
}

func doSum2(count int, val *int, waitGroup *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

func main() {
	counter := 0
	countGoroutine := 3
	var waitGroup = sync.WaitGroup{}
	var mutex = sync.Mutex{}
	
	waitGroup.Add(countGoroutine)// * 2
	for i := 0; i < countGoroutine; i++ {// * 2
		go doSum(5000, &counter, &waitGroup, &mutex)
		go doSum2(5000, &counter, &waitGroup)
	}
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}
	