package main

import (
	"fmt"
	"log"
	//"errors"
	"github.com/pkg/errors"
)




func main() {

	err := errors.New("text")
	fmt.Println(err.Error())
	fmt.Println(errors.Errorf("error writing to file, %s", "123"))
	var m interface{}
	m = my{s: "123"}

	if v, ok := m.(y); ok {
		fmt.Println(v)
	}

	s, err := sayHello("")
	if err != nil {
		log.SetFlags(0)
		log.Fatal("unexpected error is ", err)
	}
	fmt.Println(s)
}