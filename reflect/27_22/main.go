package main

import (
	"fmt"
	"reflect"
)

/*
Функции для создания новых значений

New(type) - Эта функция создает Value, указывающее на значение указанного
типа, инициализированное нулевым значением типа.

Zero(type) - Эта функция создает Value, представляющее нулевое значение
указанного типа.

MakeMap(type) - Эта функция создает новую карту, как описано в главе 28.

MakeMapWithSize(type, size) - Эта функция создает новую карту заданного размера, как описано
в главе 28.

MakeSlice(type, capacity) - Эта функция создает новый срез, как описано в главе 28.

MakeFunc(type, args, results) - Эта функция создает новую функцию с указанными аргументами и
результатами, как описано в главе 29.

MakeChan(type, buffer) - Эта функция создает новый канал с указанным размером буфера,
как описано в главе 29.

*/

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func swap(first, second interface{}) {
	firstValue, secondValue := reflect.ValueOf(first), reflect.ValueOf(second)

	if firstValue.Type() == secondValue.Type() &&
		firstValue.Kind() == reflect.Ptr &&
		firstValue.Elem().CanSet() &&
		secondValue.Elem().CanSet() {
			temp := reflect.New(firstValue.Elem().Type())
			temp.Elem().Set(firstValue.Elem())
			firstValue.Elem().Set(secondValue.Elem())
			secondValue.Elem().Set(temp.Elem())
	}
}

func main() {
	name := "Alice"
	price := 279
	city := "London"

	swap(&name, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}