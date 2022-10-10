package main

// Работа с типами массивов и срезов

import (
	"fmt"
	"reflect"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
Методы Type для массивов и срезов
Elem() - Этот метод возвращает Type для элементов массива или среза.
Len() - Этот метод возвращает длину для типа массива. Этот метод вызовет панику, если будет
вызван для других типов, включая срезы.

Функции reflect для создания массивов и типов срезов
ArrayOf(len, type) - Эта функция возвращает Type, описывающий массив с указанным размером и
 типом элемента.
SliceOf(type) - Эта функция возвращает Type, описывающий срез с указанным типом
элемента.

*/

func checkElemType(val, arrOrSlice interface{}) bool {
	elemType := reflect.TypeOf(val)
	arrOrSliceType := reflect.TypeOf(arrOrSlice)

	return (arrOrSliceType.Kind() == reflect.Array ||
		arrOrSliceType.Kind() == reflect.Slice) &&
		arrOrSliceType.Elem() == elemType
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}

	Printfln("Slice (string): %v", checkElemType("testString", slice))
	Printfln("Array (string): %v", checkElemType("testString", array))
	Printfln("Array (int): %v", checkElemType(10, array))

	sliceString := reflect.SliceOf(reflect.TypeOf(name))

	if sliceString.Kind() == reflect.Slice {

		switch sliceString.Elem().Kind() {
		case reflect.String:
			fmt.Println("Slice of string")
		}
	}
}
