package main

import (
	"fmt"
)

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