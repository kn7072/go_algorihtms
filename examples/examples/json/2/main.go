package main

import (
	"strings"
	"fmt"
	"encoding/json"
)

type DiscountedProduct struct {
	*Product `json:"product"`
	Discount float64
}
	
func main() {
	names := []string {"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int { 10, 20, 30}
	m := map[string]float64 {
		"Kayak": 279,
		"Lifejacket": 49.95,
		}
		

	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	
	dp := DiscountedProduct {
		Product: &Kayak,
		Discount: 10.50,
	}
	encoder.Encode(&dp)
		
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)
	encoder.Encode(m)
	encoder.Encode(&Kayak)


	fmt.Print(writer.String())
}
	
