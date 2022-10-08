package main

import (
	"reflect"
	//"strings"
	"fmt"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func contains(slice interface{}, target interface{}) (found bool) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() == reflect.Slice {
		for i := 0; i < sliceValue.Len(); i++ {
			if sliceValue.Index(i).Interface() == target {
				found = true
			}
		}
	}

	return
}