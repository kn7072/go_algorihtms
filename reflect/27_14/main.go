package main

import (
	"reflect"
	"strings"
	"fmt"
)

/*
Таблица 27-8 Методы Value для установки значений

CanSet() Этот метод возвращает true, если значение может быть установлено, и
false в противном случае.
SetBool(val) Этот метод устанавливает базовое значение в указанное логическое
значение.
SetBytes(slice) Этот метод устанавливает базовое значение для указанного байтового
среза.
SetFloat(val) Этот метод устанавливает базовое значение в указанное значение
float64.
SetInt(val) Этот метод устанавливает базовое значение в указанное значение int64.
SetUint(val) Этот метод устанавливает базовое значение для указанного uint64.
SetString(val) Этот метод устанавливает базовое значение в указанную строку.
Set(val) Этот метод устанавливает базовое значение в базовое значение
*/

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func incrementOrUpper(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		
		if elemValue.Kind() == reflect.Ptr {
			elemValue = elemValue.Elem()
		}
		if (elemValue.CanSet()) {
			switch (elemValue.Kind()) {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}

			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(),	elemValue)
		}
	}
}

func setAll(str interface{}, targets ...interface{}) {
	srcVal := reflect.ValueOf(str)

	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if targetVal.Kind() == reflect.Ptr &&
			targetVal.Elem().Type() == srcVal.Type() &&
			targetVal.Elem().CanSet() {
				targetVal.Elem().Set(srcVal)
		}
	}
}

func main() {
	name := "Alice"
	price := 279
	city := "London"

	incrementOrUpper(name, price, city)

	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

	incrementOrUpper(&name, &price, &city)

	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

	setAll("New String", &name, &price, &city)
	setAll(10, &name, &price, &city)

	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}
