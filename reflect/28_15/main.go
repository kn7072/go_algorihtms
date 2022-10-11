package main

// Работа с типами структур

import (
	"fmt"
	"reflect"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
Методы Type для структур

NumField() - Этот метод возвращает количество полей, определенных типом структуры.

Field(index) - Этот метод возвращает поле по указанному индексу, представленному
StructField.

FieldByIndex(indices) - Этот метод принимает срез int, который используется для поиска
вложенного поля, представленного StructField.

FieldByName(name) - Этот метод возвращает поле с указанным именем, которое представлено
StructField. Результатом является StructField, представляющий поле, и
bool значение, указывающее, было ли найдено совпадение.

FieldByNameFunc(func) - Этот метод передает имя каждого поля, включая вложенные поля, в
указанную функцию и возвращает первое поле, для которого функция
возвращает значение true. Результатом является StructField,
представляющий поле, и bool значение, указывающее, было ли найдено совпадение.

##################################################
Поля StructField
Name - В этом поле хранится имя отраженного поля.

PkgPath - Это поле возвращает имя пакета, которое используется для определения того, было ли
поле экспортировано. Для экспортируемых отраженных полей это поле возвращает
пустую строку. Для отраженных полей, которые не были экспортированы, это поле
возвращает имя пакета, который является единственным пакетом, в котором можно
использовать это поле.

Type - Это поле возвращает отраженный тип отраженного поля, описанный с помощью Type.

Tag - Это поле возвращает тег структуры, связанный с отраженным полем, как описано в
разделе «Проверка тегов структуры».

Index - Это поле возвращает int срез, обозначающий индекс поля, используемый методом
FieldByIndex, описанным в таблице 28-13.

Anonymous - Это поле возвращает значение true, если отраженное поле встроено, и значение false в
противном случае.

*/

func inspectStructType(structType reflect.Type) {
	Printfln("--- Struct Type: %v", structType)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		Printfln("Field %v: Name: %v, Type: %v, Exported: %v",
			field.Index, field.Name, field.Type, field.PkgPath == "")
	}
	Printfln("--- End Struct Type: %v", structType)
}

func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)

		if structType.Kind() == reflect.Struct {
			inspectStructType(structType)
		}
	}
}

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

func main() {
	inspectStructs(Purchase{})
}
