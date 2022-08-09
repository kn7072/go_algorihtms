package main

import (
	"fmt"
)

func Printlf(template string, args ...interface{}) {
	fmt.Printf(template + "\n", args...)
}