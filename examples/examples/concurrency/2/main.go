package main

import (
	"sync"
	"time"
)

type Message struct{
	Message int
}

type SliceMessage []Message

var channel = make(chan Message, 15)

func reader(channel <-chan Message, wg *sync.WaitGroup) {
	for val:= range channel {
		Printlf("reader %v", val)
		time.Sleep(time.Second * 2)
	}
	Printlf("Reading is stopped")
	wg.Done()
}

func readerSelect(channel <-chan Message, wg *sync.WaitGroup) {
	for {
		select {
		case val, ok := <-channel:
			Printlf("val %v, ok=%v", val.Message, ok)
			time.Sleep(time.Second * 2)
			if (!ok) {
				Printlf("Channel is closed val=%v", val.Message)
				goto exit
			}
		default:
			time.Sleep(time.Second * 2)
		}
	}
	exit:
	wg.Done()
}

func writter(channel chan<- Message, countmessage int, wg *sync.WaitGroup) {
	for i := 0; i < countmessage; i++ {
		channel <- Message{i}
		Printlf("writter %v", i)
		time.Sleep(time.Millisecond * 200) //Second
	}
	Printlf("Writting is stopped")
	close(channel)
	wg.Done()
}

func main() {
	Printlf("Start")

	count := 2 //2
	waitgroup := &sync.WaitGroup{}
	waitgroup.Add(count)

	go writter(channel, 20, waitgroup)
	//go reader(channel, waitgroup)
	go readerSelect(channel, waitgroup)

	waitgroup.Wait()
}