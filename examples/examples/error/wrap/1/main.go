package main

import (
	er "errors"
	"fmt"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err0 := errors.New("first")
	err1 := errors.Wrap(err0, "second")
	err2 := errors.Wrap(err1, "thied")
	err3 := errors.Wrap(err2, "fourth")

	err, ok := errors.Cause(err3).(stackTracer)
	if !ok {
		panic("err does not implement stackTracer")
	}
	st := err.StackTrace()
	fmt.Printf("%+v\n", st[0:3]) // top two frames

	b := er.Is(err3, err0)
	if b {
		fmt.Println("Поиск вложенных ошибок работает")
	}

	b = errors.Is(err3, err0)
	if b {
		fmt.Println("Поиск вложенных ошибок работает")
	}

	wr := errors.Unwrap(err3)
	if wr != nil {
		fmt.Printf("%+v\n", wr)
	}

	wr = er.Unwrap(err3)
	if wr != nil {
		fmt.Println(wr.Error())
		fmt.Println(er.Unwrap(wr))
	}
}
