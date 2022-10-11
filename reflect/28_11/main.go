package main

// Работа со значениями карты
// Установка и удаление значений карты
// Создание новых карт

import (
	"fmt"
	"reflect"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
Методы Value для работы с картами

MapKeys() - Этот метод возвращает значение []Value, содержащее ключи карты.

MapIndex(key) - Этот метод возвращает Value, соответствующее указанному ключу, которое
также выражается как Value. Нулевое значение возвращается, если карта не
содержит указанного ключа, что можно обнаружить, вызвав метод IsValid,
который вернет false, как описано в главе 27.

MapRange() - Этот метод возвращает *MapIter, который позволяет повторять содержимое
карты, как описано после таблицы.

SetMapIndex(key, val) - Этот метод устанавливает указанный ключ и значение, оба из которых
выражаются с использованием интерфейса Value.

Len() - Этот метод возвращает количество пар ключ-значение, содержащихся в карте.

###################################################################################

Методы, определенные структурой MapIter

Next() - Этот метод переходит к следующей паре ключ-значение на карте. Результатом этого
метода является логическое значение, указывающее, есть ли еще пары ключ-значение для
чтения. Этот метод должен вызываться перед методом Key или Value.

Key() - Этот метод возвращает Value, представляющее ключ карты в текущей позиции.

Value() - Этот метод возвращает Value, представляющее значение карты в текущей позиции.

Структура MapIter обеспечивает основанный на курсоре подход к
перечислению карт, где метод Next перемещается по содержимому карты, а
методы Key и Value обеспечивают доступ к ключу и значению в текущей
позиции. Результат метода Next указывает, есть ли оставшиеся значения для
чтения, что делает его удобным для использования с циклом for, как
показано в листинге 28-12.

###################################################################################

Функции для создания карт

MakeMap(type) - Эта функция возвращает Value, которое отражает карту, созданную с
указанным Type.

MakeMapWithSize(type, size) - Эта функция возвращает Value, которое отражает карту, созданную с
 указанным Type и размером.
*/

func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)

	if mapValue.Kind() == reflect.Map {
		for _, keyVal := range mapValue.MapKeys() {
			reflectedVal := mapValue.MapIndex(keyVal)
			Printfln("Map Key: %v, Value: %v", keyVal, reflectedVal)
		}
	} else {
		Printfln("Not a map")
	}
}

//---------------------------------------------------------

func printMapContentsRange(m interface{}) {
	mapValue := reflect.ValueOf(m)

	if mapValue.Kind() == reflect.Map {
		iter := mapValue.MapRange()
		for iter.Next() {
			Printfln("Map Key: %v, Value: %v", iter.Key(), iter.Value())
		}
	} else {
		Printfln("Not a map")
	}
}

//---------------------------------------------------------

func setMap(m, key, val interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	valValue := reflect.ValueOf(val)

	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() &&
		mapValue.Type().Elem() == valValue.Type() {
		mapValue.SetMapIndex(keyValue, valValue)
	} else {
		Printfln("Not a map or mismatched types")
	}
}

func removeFromMap(m, key interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)

	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() {
		mapValue.SetMapIndex(keyValue, reflect.Value{})
	}
}

//---------------------------------------------------------

func createMap(slice interface{}, op func(interface{}) interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)

	if sliceVal.Kind() == reflect.Slice {
		mapType := reflect.MapOf(sliceVal.Type().Elem(), sliceVal.Type().Elem())
		mapVal := reflect.MakeMap(mapType)

		for i := 0; i < sliceVal.Len(); i++ {
			elemVal := sliceVal.Index(i)
			mapVal.SetMapIndex(elemVal, reflect.ValueOf(op(elemVal.Interface())))
		}

		return mapVal.Interface()
	}

	return nil
}

func main() {
	pricesMap := map[string]float64{"Kayak": 279,
		"Lifejacket":  48.95,
		"Soccer Ball": 19.50,
	}
	printMapContents(pricesMap)

	//---------------------------------------------------------
	printMapContentsRange(pricesMap)

	//---------------------------------------------------------
	setMap(pricesMap, "Kayak", 100.00)
	setMap(pricesMap, "Hat", 10.00)
	removeFromMap(pricesMap, "Lifejacket")

	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}
	//---------------------------------------------------------
	names := []string{"Alice", "Bob", "Charlie"}
	reverse := func(val interface{}) interface{} {
		if str, ok := val.(string); ok {
			return strings.ToUpper(str)
		}

		return val
	}
	namesMap := createMap(names, reverse).(map[string]string)

	for k, v := range namesMap {
		Printfln("Key: %v, Value:%v", k, v)
	}
}
