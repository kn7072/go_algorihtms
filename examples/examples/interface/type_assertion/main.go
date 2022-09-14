package main

import (
	"fmt"
	//"github.com/pkg/errors"
)

type Child interface {
	MethodChild() int
	NewMethodChild() int
}

type Parent interface {
	Method() bool
	Child
}

type MyStuct struct {
	Value int
}

func (m *MyStuct) Method() bool {
	fmt.Printf("Method %v\n", m.Value)
	return true
}

func (m *MyStuct) MethodChild() int {
	fmt.Printf("MethodChild %v\n", m.Value)
	return m.Value
}

func (m *MyStuct) NewMethodChild() int {
	fmt.Printf("NewMethodChild %v\n", m.Value)
	return m.Value
}

func main() {
	var interfaceParent Parent
	//var interfaceChild Child
	interfaceParent = &MyStuct{Value: 1}
	//interfaceChild = interfaceParent

	if v, ok := interfaceParent.(interface{ MethodChild() int }); ok {
		fmt.Printf("%T, %v\n", v, v)
		v.MethodChild()
		//fmt.Println(v.Value)
	}

	if v, ok := interfaceParent.(*MyStuct); ok {
		fmt.Printf("%T, %v\n", v, v)
		v.MethodChild()
		fmt.Println(v.Value)
	}

	fmt.Println(interfaceParent) // , interfaceChild
}
