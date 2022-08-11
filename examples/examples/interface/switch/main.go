package main

import (
	"fmt"
)

type Interface interface {
	Method()
}

type A struct {
	ValueA int
}

type B struct {
	ValueB int
}

type C struct {
	ValueC int
}

func (a *A) Method() {
	fmt.Println(a.ValueA)
}

func (a B) Method() {
	fmt.Println(a.ValueB)
}

func (a *C) Method() {
	fmt.Println(a.ValueC)
}

func main() {
	temp := []Interface{
		&A{ValueA: 1},
		&A{},
		B{ValueB: 1},
		&C{ValueC: 1},
	}

	for _, v := range temp {
		switch value := v.(type){
		case *A:
			fmt.Println(value.ValueA)
		case B:
			fmt.Println(value.ValueB)
		case *C:
			fmt.Println(value.ValueC)
		// case C: // тип C не реализует интерфейс Interface
		// 	fmt.Println(value.ValueC)
		default:
			fmt.Println("DEFAULT", value)
		}
	}


	for _, v := range temp {
		switch value := v.(type){
		case *A, B:
			fmt.Println(value)
			if v, ok := value.(*A); ok {
				fmt.Printf("type of *A %T %v\n", v, v.ValueA)
			}
			if v, ok := value.(B); ok {
				fmt.Printf("type of B %T %v\n", v, v.ValueB)
			}
		}
	}
}