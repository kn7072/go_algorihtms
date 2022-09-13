package main

import (
	"strings"
	//"fmt"
	"encoding/json"
	"io"
)

func main() {
	x := `true "Hello" 99.99 200`
	Printfln("%T", x)
	reader := strings.NewReader(x)
	//reader := strings.NewReader("true Hello 99.99 200")
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if (err != nil) {
			if (err != io.EOF) {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
		}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
}