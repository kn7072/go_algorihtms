package main

import (
	"fmt"
)

type x interface {
	method() string
}

type y interface {
	method2() string
}

type my struct {
	s string
}

func (m *my) method() string {
	return m.s
}

func (m *my) method2() string {
	return m.s
}

type MyError struct {
	Code int
	Msg string
}

func (m *MyError) Error() string {
	return fmt.Sprintf("%s: %d", m.Msg, m.Code)
}

func sayHello(name string) (string, error) {
	if name == "" {
		return "", &MyError{Code: 2002, Msg: "no name passed"}
	}
	return fmt.Sprintf("Hello, %s", name), nil
}