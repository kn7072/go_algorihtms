package main

import "fmt"

var f = func(x int) {}

func Bar() {
	f := func(x int) {
		if x >= 0 {
			print(x)
			fmt.Printf("if Bar %T %v\n", f, f)
			f(x - 1)
		}
	}

	fmt.Printf("Bar %T %v\n", f, f)

	f(2)
}

func Foo() {
	f = func(x int) {
		if x >= 0 {
			print(x)
			fmt.Printf("if Foo %T %v\n", f, f)
			f(x - 1)
		}
	}

	fmt.Printf("Foo %T %v\n", f, f)
	f(2)
}

var b int = 1

func Test() {
	// a := 10
	ft := func() {
		print(b)
	}
	
	b := 5
	print(b)

	ft()
}

func main() {
	fmt.Printf("main %T %v\n", f, f)

	Bar()
	print(" | ")
	Foo()

	Test()
}