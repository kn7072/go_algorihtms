package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intPtr *int
	var intPtr2 *int
	temp := 10

	intPtr2 = &(temp)
	intPtrType := reflect.TypeOf(intPtr)   // array, slice, map, pointer, chan
	intPtrValue := reflect.ValueOf(intPtr) // interface, pointer

	intPtr2Type := reflect.TypeOf(intPtr2)
	intPtr2Value := reflect.ValueOf(intPtr2)

	fmt.Printf("intPtr %v\n", intPtrType.Elem()) // int

	if intPtr == nil {
		fmt.Printf("intPtr ZERO VALUE %v\n", intPtrValue.Elem())
	} else {
		fmt.Printf("intPtr %v\n", intPtrValue.Elem())
	}

	fmt.Printf("intPtr2Type %v\n", intPtr2Type.Elem()) // int

	if intPtr2 == nil {
		fmt.Printf("intPtr2 ZERO VALUE %v\n", intPtr2Value.Elem())
	} else {
		fmt.Printf("intPtr2 %v\n", intPtr2Value.Elem())
	}
}
