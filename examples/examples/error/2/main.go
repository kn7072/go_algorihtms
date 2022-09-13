package main

import (
	//"errors"
	"fmt"
	"github.com/pkg/errors"
)

type structer struct {
	a []int
}

func (s *structer) Change() error{
	s.a = append(s.a, 100)
	//err := errors.New("Error Change") //nil
	err := fmt.Errorf("My error code is %d\n", 32)
	return err
}	

func (s structer) Change2() {
	s.a = append(s.a, 200)
}

func inner() error {
	a := structer{
		a: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(a)
	err := a.Change()
	if err != nil {
		return fmt.Errorf("inner err %v\n", err)
	}
	fmt.Println(a)
	return nil
}

func main() {
	
	err := inner()
	if err != nil {
		fmt.Println(err)
	}

	err1 := errors.New("ERROR")
	err2 := errors.Wrap(err1, "open failed")
	err3 := errors.Wrap(err2, "read config failed")
	fmt.Println(err3)
	fmt.Printf("%+v", err3) // напечатает stacktrace

	// Cause доходит до первопричины ошибки - err1
	print(err1 == errors.Cause(err3)) // true проверка вложенности
	print(err1 == errors.Cause(err2)) // true
	
    fmt.Printf("%+v", errors.Cause(err3))
	print(err2 == errors.Cause(err3)) // false
}