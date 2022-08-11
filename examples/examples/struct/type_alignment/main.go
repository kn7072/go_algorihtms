package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type someData struct{
	a int8  // 1 byte
	b int64 // 8 byte
	c int8  // 1 byte
 }

type someDataV2 struct{
	a int8  // 1 byte
	c int8  // 1 byte
	b int64 // 8 byte
 }

// https://ubiklab.net/posts/go-alignment/

func main() {
	v := someData{}
	typ := reflect.TypeOf(v)
	fmt.Printf("Type someData is %d bytes long\n", typ.Size())

	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(), field.Type.Align())
	}
	fmt.Printf("someData align is %d\n", typ.Align())
	/*
		Offset — это смещение адреса поля в значении структуры. 
		Size — размер поля. 
		Align — выравнивание для типа поля.
	*/

	/*
		Почему такой результат. 
		В общем случае экземпляр структуры будет выровнен по самому длинному элементу. 
		Для компилятора это самый простой способ убедиться, что все поля структуры будут также выровнены для быстрого доступа.
		
		Поле v.a занимает всего один байт, но следующее поле начинается только через 8 байт (b at offset 8). 7 байт просто не используется.
	*/

	//Посмотрим подробней как это выглядит в памяти:
	v = someData{a:1,b:2,c:3}
	b := (*[24]byte)(unsafe.Pointer(&v))
	fmt.Printf("Bytes are %#v\n", b)

	/*
		Компилятор не меняет порядок полей в структуре и не может оптимизировать такие случае. А мы можем.
		Смежные поля могут быть объедены, если их сумма не превышает выравнивания структуры.
	*/

	fmt.Println("#######################################")
	v2 := someDataV2{}
	typ2 := reflect.TypeOf(v2)
	fmt.Printf("Type someData is %d bytes long\n", typ2.Size())
	
	n2 := typ2.NumField()
	for i := 0; i < n2; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(), field.Type.Align())
	}
	fmt.Printf("someDataV2 align is %d\n", typ.Align())

	v2 = someDataV2{a:1,b:2,c:3}
	b = (*[24]byte)(unsafe.Pointer(&v))
	fmt.Printf("Bytes are %#v\n", b)

	/*
		Почти всегда, гораздо важней читаемость кода, чем такие оптимизации. 
		Нужно понимать по какой причины у значения типа именно такой размер и что вообще происходит, а уже в случае необходимости заниматься оптимизацией.

		Линтеры, которые могут помочь:

			https://gitlab.com/opennota/check
			https://github.com/mdempsky/maligned

	*/
}