package main

// Работа со значениями массива и среза
// Перечисление срезов и массивов

import (
	"fmt"
	"reflect"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
Value методы работы с массивами и срезами

Index(index) - Этот метод возвращает Value, представляющее элемент по указанному индексу.
Len() - Этот метод возвращает массив или длину среза.
Cap() - Этот метод возвращает емкость массива или среза.
SetLen() - Этот метод устанавливает длину среза. Его нельзя использовать в массивах.
SetCap() - Этот метод устанавливает емкость среза. Его нельзя использовать в массивах.
Slice(lo, hi) - Этот метод создает новый срез с указанными нижним и верхним значениями.
Slice3(lo, hi, max) - Этот метод создает новый срез с указанными минимальными, высокими и 
максимальными значениями.
 
*/

func setValue(arrayOrSlice interface{}, index int, replacement interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacement)

	if (arrayOrSliceVal.Kind() == reflect.Slice) {
		elemVal := arrayOrSliceVal.Index(index)
		if (elemVal.CanSet()) {
			elemVal.Set(replacementVal)
		}
	} else if (arrayOrSliceVal.Kind() == reflect.Ptr &&
		arrayOrSliceVal.Elem().Kind() == reflect.Array &&
		arrayOrSliceVal.Elem().CanSet()) {
			arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}
}

func enumerateStrings(arrayOrSlice interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)

	if (arrayOrSliceVal.Kind() == reflect.Array ||
		arrayOrSliceVal.Kind() == reflect.Slice) &&
		arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
			for i := 0; i < arrayOrSliceVal.Len(); i++ {
				Printfln("Element: %v, Value: %v", i, arrayOrSliceVal.Index(i).String())
			}
	}
}

func findAndSplit(slice, target interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)

	if (sliceVal.Kind() == reflect.Slice && 
		sliceVal.Type().Elem() == targetType) {
			for i := 0; i < sliceVal.Len(); i++ {
				if sliceVal.Index(i).Interface() == target {
					return sliceVal.Slice(0, i +1)
				}
			}
	}

	return slice
}


func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}

	Printfln("Original slice: %v", slice)
	newCity := "Paris"
	setValue(slice, 1, newCity)
	Printfln("Modified slice: %v", slice)
	Printfln("Original slice: %v", array)

	newCity = "Rome"
	setValue(&array, 1, newCity)
	Printfln("Modified slice: %v", array)

	//-----------------
	enumerateStrings(slice)
	enumerateStrings(array)

	//-----------------
	name2 := "Alice"
	city2 := "London"
	hobby2 := "Running"
	slice2 := []string {name2, city2, hobby2}
	Printfln("Strings: %v", findAndSplit(slice2, "London"))

	numbers := []int {1, 3, 4, 5, 7}
	Printfln("Numbers: %v", findAndSplit(numbers, 4))
}
	