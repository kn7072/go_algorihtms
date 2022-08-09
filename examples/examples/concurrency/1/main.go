package main

import (
	"time"
	"runtime"
)

type Message struct{
	Message string
}

type sliceMessage []Message

var chanNB = make(chan Message)
var chanB = make(chan Message, 10)

func readerNB(channel <-chan Message){
	
	for {
		select {
		case val, ok := <- channel:
			//Printlf("reader %v", <-channel)
			if (!ok) {
				// канал закрыт - выходим
				goto exit
			}
			Printlf("reader %v, channel is open %v", val, ok)
			Printlf("___________________")
		default:
			time.Sleep(time.Second * 1)	
		}
	}
	exit:
	Printlf("Readings is stoped")
}

func writterNB(channel chan<- Message, slicemessage sliceMessage){
	for _, val := range slicemessage {
		Printlf("writter v %v, T %T", val.Message , val)
		channel <- val//.Message
		time.Sleep(time.Second * 2) // Microsecond
	}
	close(channel)
}

var const1 int = 1

func testDefer(i int) {
	Printlf("TEST DEFER i %v, %v", i, const1)
}

func makeRecover() {
	defer func () {
		panicValue := recover()
		Printlf("panic - %v", panicValue)
	}()

	panic("some panic")
}

func main() {
	i := 0
	defer makeRecover()
	defer testDefer(i)
	Printlf("Start")
	//runtime.GOMAXPROCS(2)
	Printlf("NumCPU %v", runtime.NumCPU())
	time.Sleep(time.Second)

	testSliceMessage := sliceMessage{
		{Message: "one"},
		{Message: "two"},
		{Message: "three"},
		{Message: "four"},
	}
	go readerNB(chanNB)
	go writterNB(chanNB, testSliceMessage)
	time.Sleep(time.Second * 20)
	i = 10
	const1 = 10
}
