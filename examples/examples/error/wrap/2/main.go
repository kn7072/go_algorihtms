package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
)

type MyError struct {
	Value int
}

type MyErrorInterface interface {
	Method() bool
}

func (m *MyError) MyErrorInterface() bool {
	return true
}

type MyError2 struct {
	Value int
}

func TestTypeStruct(t MyError) {
	fmt.Println(t.Value)
}

func TestTypeStruct2(t struct{ Value int }) {
	fmt.Println(t.Value)
}

func (e *MyError) Error() string {
	return fmt.Sprintf("My error %v", e.Value)
}

var (
	errorTemp0 = errors.New("ERROR CONST")
	errorTemp  = &MyError{Value: 100}
	wrapError  = fmt.Errorf("WRAP %w", errorTemp)
	//wrapError1  = fmt.Errorf("WRAP %w", errorTemp)
)

func TestRecover(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println(err)
			}
		}
	}()

	err := fmt.Errorf("Test panic")
	panic(err)
	//wg.Done()
}

func runValidation() error {
	return fmt.Errorf("WRAP error %w", wrapError)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go TestRecover(wg)
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())

	err := runValidation()
	if err == wrapError || errors.Unwrap(err) == wrapError {
		fmt.Println("Ошибка ", err)
	} else {
		fmt.Println("-")
	}

	if err == errorTemp || errors.Unwrap(err) == errorTemp {
		fmt.Println("Ошибка ", err)
	} else if errors.Is(err, errorTemp) {
		fmt.Println("Проверка на вложенность")
	}

	var valueError *MyError
	//var valueError *MyError
	if b := errors.As(err, &valueError); b {
		fmt.Println(valueError.Value)
	}

	errorTest := error(&MyError{Value: 100})
	if x, ok := errorTest.(interface{ Method() bool }); ok { // && x.Is(valueError)
		fmt.Printf("%T %v\n", x, x)
	}
	if x, ok := errorTest.(*MyError); ok {
		fmt.Printf("%T %v\n", x, x)
	}

	//#####################################33

	mError1 := MyError{Value: 1}
	mError2 := MyError2{Value: 1}
	TestTypeStruct(mError1)
	//TestTypeStruct(mError2)

	TestTypeStruct2(mError1)
	TestTypeStruct2(mError2)

	if mError1 == MyError(mError2) {
		fmt.Println("mError1 = mError2")
	}

	// if MyError == MyError2 {
	// 	fmt.Println("mError1 = mError2")
	// }
}
