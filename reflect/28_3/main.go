package main

// Работа с указателями
// Работа со значениями указателя

import (
	"fmt"
	"reflect"
	"strings"
)

type Product struct {
	Name, Category string
	Price          float64
}

type Customer struct {
	Name, City string
}

type Purchase struct {
	Customer
	Product
	Total   float64
	taxRate float64
}

/*
Функция пакета reflect и метод для указателей

PtrTo(type) - Эта функция возвращает Type, который является указателем на Type, полученный в
качестве аргумента.

Elem() - Этот метод, который вызывается для типа указателя, возвращает базовый Type. Этот
метод вызывает панику при использовании для типов, не являющихся указателями.

Функция PtrTo создает тип указателя, а метод Elem возвращает тип, на
который указывает указатель, как показано в листинге
*/

/*
Методы Value для работы с типами указателей

Addr() - Этот метод возвращает Value, которое является указателем на Value, для которого он
вызывается. Этот метод вызывает панику, если метод CanAddr возвращает значение
false.

CanAddr() - Этот метод возвращает значение true, если значение можно использовать с методом
Addr.

Elem() - Этот метод следует за указателем и возвращает его Value. Этот метод вызывает панику,
если он вызывается для значения, не являющегося указателем.

*/

func createPointerType(t reflect.Type) reflect.Type {
	return reflect.PtrTo(t)
}

func followPointerType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}

	return t
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

var stringPtrType = reflect.TypeOf((*string)(nil))

func transformString(val interface{}) {
	elemValue := reflect.ValueOf(val)

	if elemValue.Type() == stringPtrType {
		upperStr := strings.ToUpper(elemValue.Elem().String())
		if elemValue.Elem().CanSet() {
			elemValue.Elem().SetString(upperStr)
		}
	}
}

func main() {
	name := "Alice"
	t := reflect.TypeOf(name)
	Printfln("Original Type: %v, %T", t, t)
	pt := createPointerType(t)
	Printfln("Pointer Type: %v, %T", pt, pt)
	Printfln("Follow pointer type: %v", followPointerType(pt))

	//-------------------------------------------------------

	transformString(&name)
	Printfln("Follow pointer value: %v", name)
}