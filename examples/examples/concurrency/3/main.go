package main

// Mutex

import (
	"time"
	"sync"
	//"context"
)

type Message struct{
	Message string
}

type SliceMessages []Message
var sliceM = make(SliceMessages, 1, 20)

var mutex = sync.Mutex{}

func init() {
	temp := []string{"a", "b", "c"}
	for _, val := range temp{
		sliceM = append(sliceM, Message{val})
	}
}

func reader(channel <-chan Message, wg *sync.WaitGroup, 
	        sliceM *SliceMessages, mutex *sync.Mutex){
	for {
		select {
		case _, ok := <-channel:
			if (!ok){
				Printlf("Channel is closed")
				goto exit
			}
		default:
			if (len(*sliceM) == 0) {
				time.Sleep(time.Second * 1)
			} else {
				mutex.Lock()
				Printlf("last value is %v", (*sliceM)[len(*sliceM) - 1])
				mutex.Unlock()
				time.Sleep(time.Second * 1)
			}
		}
	}
	exit:
	wg.Done()
}

func writter(channel chan<- Message, wg *sync.WaitGroup, 
	         sliceM *SliceMessages, mutex *sync.Mutex) {
		temp := []string{"D", "E", "F", "G", "J"}
		time.Sleep(time.Second * 3)
		for _, val := range temp{
			mutex.Lock()
			(*sliceM)[len(*sliceM)-1] = Message{val}
			mutex.Unlock()
			time.Sleep(time.Second * 5)
		}
		close(channel)
		wg.Done()
	}


func main(){
	waitGroup := &sync.WaitGroup{}
	var channel = make(chan Message)

	waitGroup.Add(2)
	Printlf("sliceM %v", sliceM)
	go reader(channel, waitGroup, &sliceM, &mutex)
	go writter(channel, waitGroup, &sliceM, &mutex)
	waitGroup.Wait()
}