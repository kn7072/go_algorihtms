package main

import "fmt"

type MyInterface interface {
	Method()
}

type MyInterfaceTypes interface {
	int | int32 | int64
}

type Test struct {
	test string
}

func (t *Test) Method() {
	fmt.Printf("Method t.test %s\n", t.test)
}

func Reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	for _, el := range list {
		init = accumulator(el, init)
	}
	return init
}

func Ternary[T any](cond bool, x T, y T) T {
	if cond {
		return x
	}
	return y
}

func main() {
	sStr := "hallo"
	ssStr := []string{"a", "b", "c"}

	ssInt := []int{1, 2, 3, 4, 5} 

	fmt.Println(ContainsSrt(sStr, ssStr))

	sum := func(a, b string) string {return a + b}
	fmt.Println(Reduce(ssStr, sum, ""))

	mul := func(a, b int) int { return a * b }
	fmt.Println(Reduce(ssInt, mul, 1))

	fmt.Println(Ternary(5 > 2, "true", "false"))

	sInt := 1
	
	fmt.Println(ContainsInt(sInt, ssInt))

	fmt.Println(Contains(sInt, ssInt))

	PrintSlice(ssStr)
	PrintSlice(ssInt)

	testSliceMyinterface := []MyInterface{&Test{"test_generic"}, 
								          &Test{"test_generic_2"},
							}
	Print(testSliceMyinterface)

	strs := []*Test{{"test_generic_1"}, 
					{"test_generic_2"}, 
					{"test_generic_3"}, 
				}
	Print(strs)
}

func Print[T MyInterface](s []T) {
	for _, v := range s {
		fmt.Println(v)
		v.Method()
	}

}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Contains[T comparable](v T, vv []T) bool {
	for _, a := range vv {
		if a == v {
			return true
		}
	}
	return false
}

func ContainsSrt(v string, vv []string) bool {
	for _, a := range vv {
		if a == v {
			return true
		}
	}
	return false
}

func ContainsInt(v int, vv []int) bool {
	for _, a := range vv {
		if a == v {
			return true
		}
	}
	return false
}