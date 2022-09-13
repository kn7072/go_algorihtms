package main

import (
	"strings"
	//"fmt"
	"encoding/json"
	"io"
)

func main() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	//vals := []interface{} {}
	ints := []int {}
	mixed := []interface{} {}
	vals := []interface{} { &ints, &mixed}


	decoder := json.NewDecoder(reader)
	for i:=0; i<len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			if err != io.EOF {
			Printfln("Error %v", err.Error())
			}
			break
		}
	}
	// for {
	// 	var decodedVal interface{}
	// 	err := decoder.Decode(&decodedVal)
	// 	if (err != nil) {
	// 		if (err != io.EOF) {
	// 			Printfln("Error: %v", err.Error())
	// 		}
	// 		break
	// 	}
	// 	vals = append(vals, decodedVal)
	// }
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
}