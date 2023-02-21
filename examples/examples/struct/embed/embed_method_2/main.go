package main

// https://stackoverflow.com/questions/52106242/golang-inheritance-and-method-override

import (
	"fmt"
	"sync"
)

type Parent struct {
	sync.Mutex
	MyInterface
}

type Empty struct {
	MyInterface
}

func (p *Parent) Foo() {
	p.Lock()
	defer p.Unlock()
	p.Bar()
}

func (p *Parent) B() {
	panic("NOT IMPLEMENTED")
}

func (p *Parent) A() {
	p.Lock()
	defer p.Unlock()
	p.B()
}

type MyInterface interface {
	Foo()
	Bar()
}

type Child struct {
	Parent
	Name string
}

func (c *Child) Bar() {
	fmt.Println(c.Name)
}

func (c *Child) B() {
	fmt.Println(c.Name)
}

// ----------------------------------
type Implemented struct {
	Name string
}

func (imp Implemented) Foo() {
	fmt.Println("Implemented Foo")
}

func (imp Implemented) Bar() {
	fmt.Println("Implemented Bar")
}

//----------------------------------

func main() {
	c := new(Child)
	c.Name = "Child"
	// c.A()   // panic
	// c.Foo() // invalid memory address or nil pointer dereference
	e := Empty{}
	// e.Bar() invalid memory address or nil pointer dereference
	fmt.Printf("%#v\n", e)

	eno := Empty{Implemented{}}
	eno.Bar()
}
