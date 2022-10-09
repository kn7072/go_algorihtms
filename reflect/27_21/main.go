package main

import (
	"fmt"
	"reflect"
)

/*
Методы Value для проверки переполнения

OverflowFloat(val) - Этот метод возвращает значение true, если указанное значение
float64 вызовет переполнение при преобразовании в тип Value, для
которого вызывается метод. Этот метод вызовет панику, если только
метод Value.Kind не вернет Float32 или Float64.

OverflowInt(val) - Этот метод возвращает значение true, если указанное значение int64
вызовет переполнение при преобразовании в тип Value, для которого
вызывается метод. Этот метод вызовет панику, если только метод
Value.Kind не вернет один из целочисленных типов со знаком.

OverflowUint(val) - Этот метод возвращает значение true, если указанное значение
uint64 вызовет переполнение при преобразовании в тип Value, для
которого вызывается метод. Этот метод вызовет панику, если только
метод Value.Kind не вернет один из целочисленных типов без знака.

*/

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func IsInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}

	return false
}

func IsFloat(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}

	return false
}

func convert(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)

	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		if (IsInt(targetVal) && IsInt(srcVal)) &&
			targetVal.OverflowInt(srcVal.Int()) {
			Printfln("Int overflow")
			return src, false
		} else if IsFloat(targetVal) && IsFloat(srcVal) &&
			targetVal.OverflowFloat(srcVal.Float()) {
			Printfln("Float overflow")
			return src, false
		}

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

	newVal, ok = convert(5000, int8(100))
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
}
