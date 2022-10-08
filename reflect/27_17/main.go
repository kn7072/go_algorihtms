package main

import (
	"reflect"
	//"strings"
	"fmt"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
	Comparable() - Этот метод возвращает true, если отраженный тип можно использовать с
	оператором сравнения Go, и false в противном случае.
*/

func contains(slice interface{}, target interface{}) (found bool) {
	sliceValue := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)

	if sliceValue.Kind() == reflect.Slice &&
		sliceValue.Type().Elem().Comparable() &&
		targetType.Comparable() {
		for i := 0; i < sliceValue.Len(); i++ {
			if sliceValue.Index(i).Interface() == target {
				found = true
			}
		}
	}

	return
}

func main() {
	city := "London"
	citiesSlice := []string{ "Paris", "Rome", "London"}

	Printfln("Found #1: %v", contains(citiesSlice, city))

	sliceOfSlices := [][]string{
		citiesSlice, {"First", "Second", "Third"}}
		
	Printfln("Found #2:	%v", contains(sliceOfSlices, citiesSlice))
}