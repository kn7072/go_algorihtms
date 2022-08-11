package main

import (
	"fmt"
)

/*
https://stackoverflow.com/questions/65440847/invalid-receiver-for-pointer-alias-type
https://go.dev/ref/spec#Method_declarations

The receiver is specified via an extra parameter section preceding the method name. 
That parameter section must declare a single non-variadic parameter, the receiver. 
Its type must be a defined type T or a pointer to a defined type T. 
T is called the receiver base type. 
A receiver base type cannot be a pointer or interface type and it must be defined in the same package as the method.


type A struct { 
	value int
}
type B *A

func (b B)Print() { // приемник тип указатель
	fmt.Printf("Value: %d\n", b.value)
}



// интерфейс для тестирования ресивера с типом интерфейс
type MyInterface interface {
	Method()
}

func (i MyInterface) Method() { // приемник интерфейс - не законно
	fmt.Println(i)
}
*/

type MyStuct struct {
	Value int
}

func (p *MyStuct) MethodPointer(newValue int) {
	fmt.Printf("old %v\n", p.Value)
	p.Value = newValue
}

func (p MyStuct) MethodObject(newValue int) {
	fmt.Printf("old %v\n", p.Value)
	p.Value = newValue
}


func FuncPointer(x *MyStuct, newValue int) {
	fmt.Printf("old %v\n", x.Value)
	x.Value = newValue
}

func FuncObject(x MyStuct, newValue int) {
	fmt.Printf("old %v\n", x.Value)
	x.Value = newValue
}

func main() {
	obj := MyStuct{Value: 1}
	
	obj.MethodPointer(2)
	fmt.Println(obj.Value)

	(&obj).MethodPointer(22)
	fmt.Println(obj.Value)

	obj.MethodObject(3)
	fmt.Println(obj.Value)

	(&obj).MethodObject(4)
	fmt.Println(obj.Value)

	fmt.Printf("############################")

	// Передаем указатель
	FuncPointer(&obj, 0)
	fmt.Println(obj.Value) // 0

	//FuncPointer(obj, 0) - нужен указатель а не объект
	
	FuncObject(obj, 1)
	fmt.Println(obj.Value) // 1

	//FuncObject(&obj, 2)

}