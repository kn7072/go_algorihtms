package main

import (
	"fmt"
)

var o = fmt.Print

type A int

func (A) g()     { o("A_g\n") }
func (A) m() int { return 1 }

type B int

func (B) g() { o("B_g\n") }
func (B) f() { o("B_f\n") }

type C struct {
	A
	B
}

func (C) m() int { return 9 }

func main() {
	var c interface{} = C{}
	_, bf := c.(interface{ f() })
	_, bg := c.(interface{ g() })
	i := c.(interface{ m() int })

	fmt.Println(bf, bg, i.m())

	i.(C).f()

	// i.(C).g()
	i.(C).A.g()
	i.(C).B.g()
}
