package main

import (
	"fmt"
	"reflect"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
Функции добавления элементов к срезам

MakeSlice(type, len, cap) - Эта функция создает Value, отражающее новый срез, используя Type для
обозначения типа элемента с заданной длиной и емкостью.

Append(sliceVal, ...val) - Эта функция добавляет к указанному срезу одно или несколько значений,
все из которых выражаются с помощью интерфейса Value. Результатом является измененный срез.
Функция вызывает панику, когда используется для любого типа, отличного от среза, или если типы значений не
соответствуют типу элемента среза.

AppendSlice(sliceVal, sliceVal) - Эта функция добавляет один срез к другому. Функция паникует, если либо
Value не представляет срез, либо если типы срезов несовместимы.

Copy(dst, src) - Эта функция копирует элементы из среза или массива, отраженного src
Value, в срез или массив, отраженный dst Value. Элементы копируются
до тех пор, пока целевой срез не будет заполнен или пока не будут
скопированы все исходные элементы. Источник и место назначения
должны иметь один и тот же тип элемента.

*/

func pickValues(slice interface{}, indices ...int) interface{} {
	sliceVal := reflect.ValueOf(slice)

	if sliceVal.Kind() == reflect.Slice {
		newSlice := reflect.MakeSlice(sliceVal.Type(), 0, 10)
		for _, index := range indices {
			newSlice = reflect.Append(newSlice, sliceVal.Index(index))
		}

		return newSlice
	}

	return nil
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"

	slice := []string{name, city, hobby, "Bob", "Paris", "Soccer"}
	picked := pickValues(slice, 0, 3, 5)
	Printfln("Picked values: %v", picked)
}
