package main

// https://medium.com/learning-the-go-programming-language/comparing-values-in-go-8f7b002e767a

import (
	"fmt"
	"math"
)

type shape interface {
	area() int
}

type Color interface {
	GetColor() string
}

type rectangle struct {
	l int
	w int
}

func (r rectangle) area() int {
	return r.l * r.w
}

type rectangleNew struct {
	l int
	w int
}

func (r rectangleNew) area() int {
	return r.l * r.w
}

type square struct {
	l int
}

func (s square) area() int {
	return s.l * s.l
}

/*
Comparing Values

Equality is tested using the == and the != comparison operators. 
Most types can be tested for equality, some have limited support while others,
 however, not at all. The ordering operators <, <=, >, and >= are used to test values
  with types that can be ordered. Some types can be ordered while others not all.
   Let us look at how to compare each type in Go.
*/

func main() {
	/*
	Boolean

	Booleans can be compared to pre-defined values of (or expressions that produce) true 
	or false. It is an error to attempt to compare a boolean value to an numeric value.
	*/
	a := true
	if a != (10 == 20) {
		fmt.Println("a not true")
	}
	// following blows up at compilation
	// if a == 1 {
	// 	fmt.Println("some useful")
	// }

	/*
	Integer and Floating-point numbers

	Comparing numerical values works as you would expect, following the general rule above,
	 for both equality and order.
	*/
	aNum := 3.1415
	if aNum != math.Pi {
		fmt.Println("aNum is not pi")
	}

	/*
	Complex Numbers

	Complex numbers can also be tested for equality. 
	Two complex values are equal if their real and imaginary parts are equal respectively.
	*/
	aComplex := complex(-3.25, 2)
  	bComplex := -3.25 + 2i

  	if aComplex == bComplex {
    	fmt.Println("aComplex complex as bComplex")
  	}

	/*
	Due to the nature of complex numbers, however, 
	they do not support ordering operators in Go.
	if aComplex < bComplex {
    	fmt.Println("a complex as b")
  	}
	*/

	/*
	String Values

	String values support both the standard equality and ordering operators. 
	There are no additional functions needed to compare strings. 
	Values can be automatically compared lexically using ==, !=, <=, <, >, and >= operators.
	*/
	cols := []string{
        "xanadu", "red", "fulvous", 
        "white", "green", "blue",
        "orange", "black", "almond"}
    for _, col := range cols {
        if col >= "red" || col == "black" {
          fmt.Println(col)
        }
    }

	/*
	Struct Values

	Two struct values can be tested for equality by comparing the values 
	of their individual fields. In general, two struct values are considered equal 
	if they are of the same type and the their corresponding fields are equal.
	*/

	p1 := struct {a string; b int}{"left", 4}
    p2 := struct {a string; b int}{a: "left", b: 4}

    if p1 == p2 {
        fmt.Println("Same position")
    }

	/*
	In the previous code snippet, struct p1 is equal to p2 since they are of the same type 
	and their corresponding field values are the same. 
	Any change in the field values will cause the structs to be not equal.

	Struct values, however, cannot be compared using ordering operators. 
	So the following code will not compile.
	if p1 > p2 {
        fmt.Println("Same position")
    }
	*/

	/*
	Arrays

	Array values are compared for equality by comparing elements of the their defined types.
	 Arrays are equal if their corresponding values are equal.
	*/
	pair1 := [2]int {4, 2}
    pair2 := [2]int {2, 4}

    if pair1 != pair2 {
        fmt.Println("different pair")
    }
	/*
	As with struct values, arrays cannot be compared using 
	ordering operators <, <=, >, and >=. Attempting to do so will cause a compilation error.
	*/

	/*
	Pointers

	Pointer values can be compared for equality but not for ordering. 
	Two pointer values are considered equal if they point to the same value in memory (or if they are nil). For instance, in the following snippet &pair is equal to ptr2 while &pair is not equal ptr.
	*/

	pair := [2]int{4, 2}
    ptr := &[2]int{4, 2}
    ptr2 := &pair
 
    if &pair != ptr {
        fmt.Println("pointing different")
    }

    if &pair == ptr2 {
        fmt.Println("pointing the same")
    }

	/*
	Keep in mind that a pointer to a type is not the same as the type itself. 
	Thus Trying to compare the two will cause a type mismatched compilation error.
	*/

	/*
	Interfaces

	Interface values are interesting in that they can be compared to

    - not only other interface values
    - but also to values whose types implement the interface

	Two interface values are considered equal if their underlying concrete 
	types and their values are comparable and are equal or if both interfaces are nil.
	
	For instance, in the next code snippet, interface values r0 and r2 are equal 
	because they implement the same concrete types with the same values, 
	rectangle{l:3, w:6}. On the other hand, interface values r0 and r1, 
	although they implement the same interface type, are understandably not equal 
	because their concrete values differ, rectangle{3, 6} vs rectangle{6, 3}. 
	Similarly, interface variables r1 and s0 are not equal because they have 
	different dynamic (or concrete) values, although they implement the same interface.
	*/

	var r0 shape = rectangle{3, 6}
   	var r1 shape = rectangle{6, 3}
   	var r2 shape = rectangle{3, 6}
	var r3 shape = rectangleNew{3, 6}
   	var s0 shape = square{5}
	
	var color Color
	var color2 Color
	// var shapeIntreface shape
	// if color == shapeIntreface {
	// 	fmt.Println("error")
	// }
	if color == color2 {
		fmt.Println("color and color2 same Color")
	}

   	if r0 == r2 {
    	fmt.Println("r0 and r2 same shapes")
   	}

	fmt.Println("r1 and r3 equal", (r1 == r3))
 
	fmt.Println("r1 and s0 equal", (r1 == s0))
	/*
	It is important to note that if the underlying concrete types of the interface values
	 are not comparable (see previous section on topic), any attempt to compare 
	 them will cause a runtime panic.
	*/

	/*
	Channels

	Channel values can only be compared for equality. Two channel values are considered equal 
	if they originated from the same make call 
	(meaning they refer to the same channel value in memory).

	For instance, in the following sample, 
	ch0 is not equal to ch1 even when they have the same types. 
	However, ch1 is equal to ch2 because they both refer to the same channel value.
	*/

	ch0 := make(chan int)
 	ch1 := make(chan int)
 	ch2 := ch1
 
 	fmt.Println("ch0 == ch1", (ch0==ch1))
 	fmt.Println("ch1 == ch2", (ch1==ch2))
}