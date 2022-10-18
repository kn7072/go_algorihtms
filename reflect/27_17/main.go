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

func contains(slice, target interface{}) (found bool) {
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

func containsDEBUG(slice, target interface{}) (found bool) {
	sliceValue := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)

	if sliceValue.Kind() == reflect.Slice {
		fmt.Println("sliceValue.Kind() == reflect.Slice")
		fmt.Printf("    sliceValue.Type() %v\n", sliceValue.Type())
		fmt.Printf("    sliceValue.Type().Elem() %v\n", sliceValue.Type().Elem())

		if sliceValue.Type().Elem().Comparable() {
			fmt.Println("sliceValue.Type().Elem().Comparable()")

			if targetType.Comparable() {
				fmt.Println("targetType.Comparable()")

				for i := 0; i < sliceValue.Len(); i++ {
					if sliceValue.Index(i).Interface() == target {
						found = true
					}
				}
			}
		}
	}

	return
}

func main() {
	city := "London"
	citiesSlice := []string{"Paris", "Rome", "London"}

	Printfln("Found #1: %v", contains(citiesSlice, city))
	Printfln("Found #1: %v", containsDEBUG(citiesSlice, city))

	sliceOfSlices := [][]string{
		citiesSlice, {"First", "Second", "Third"}}

	Printfln("Found #2:	%v", contains(sliceOfSlices, citiesSlice))
	Printfln("Found #2:	%v", containsDEBUG(sliceOfSlices, citiesSlice))
}
