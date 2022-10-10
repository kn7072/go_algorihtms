package main

import (
	"fmt"
	"reflect"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
Методы Type для карт
Key() - Этот метод возвращает Type для ключей карты.
Elem() -  Этот метод возвращает Type для значений карты.

Функции reflect для создания типов карт
MapOf(keyType, valType) - Эта функция возвращает новый Type, который отражает тип карты с указанными
типами ключа и значения, оба из которых описаны с использованием Type.
*/

func describeMap(m interface{}) {
	mapType := reflect.TypeOf(m)

	if mapType.Kind() == reflect.Map {
		Printfln("Key type: %v, Val type: %v", mapType.Key(), mapType.Elem())
	} else {
		Printfln("Not a map")
	}
}

func main() {
	pricesMap := map[string]float64{"Kayak": 279,
		"Lifejacket":  48.95,
		"Soccer Ball": 19.50,
	}
	describeMap(pricesMap)
}
