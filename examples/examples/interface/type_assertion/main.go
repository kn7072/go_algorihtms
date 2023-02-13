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

type I1 interface {
	M()
}

type I2 interface {
	I1
	N()
}

type T struct {
	name string
}

func (T) M() {}
func (T) N() {}

func main() {
	var interfaceParent Parent
	//var interfaceChild Child
	interfaceParent = &MyStuct{Value: 1}
	//interfaceChild = interfaceParent

	if v, ok := interfaceParent.(interface{ MethodChild() int }); ok {
		fmt.Printf("%T, %v\n", v, v)
		v.MethodChild()
		//fmt.Println(v.Value)
		if vDinamiv, ok := v.(*MyStuct); ok {
			fmt.Printf("%T, %v, %v", vDinamiv, vDinamiv, vDinamiv.Value)
		}
	}

	if v, ok := interfaceParent.(*MyStuct); ok {
		fmt.Printf("%T, %v\n", v, v)
		v.MethodChild()
		fmt.Println(v.Value)
	}

	fmt.Println(interfaceParent) // , interfaceChild

	var v1 I1 = T{"foo"}
	var v2 I2
	v2, ok := v1.(I2)
	fmt.Printf("%T %v %v\n", v2, v2, ok) // main.T {foo} true
	// fmt.Printf("v2.name", v2.name)
	v2.N()
}
