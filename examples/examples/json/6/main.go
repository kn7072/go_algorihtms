package main

import (
	"strings"
	//"fmt"
	"encoding/json"
	//"io"
)

func main() {

	reader := strings.NewReader(`true "Hello" 99.99 200`)
	var	bval bool
	var	sval string
	var	fpval float64
	var	ival int
	vals := []interface{} { &bval, &sval, &fpval, &ival }
	decoder := json.NewDecoder(reader)
	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i]) //vals[i] &bval
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T):	%v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T):	%v", fpval, fpval)
	Printfln("Decoded (%T):	%v", ival, ival)
}