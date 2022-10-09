package main

import (
	"fmt"
	"reflect"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
ConvertibleTo(type) - Этот метод возвращает значение true, если тип, для которого
вызывается метод, может быть преобразован в указанный Type.

Метод Value для преобразования типов
Convert(type) - Этот метод выполняет преобразование типа и возвращает Value с новым
типом и исходным значением.
*/

func convert(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)

	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}

	return
}

func main() {
	name := "Alice"
	price := 279

	newVal, ok := convert(price, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)

	newVal, ok = convert(name, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
}