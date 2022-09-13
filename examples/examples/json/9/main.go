package main

import (
	"strings"
	//"fmt"
	"encoding/json"
	"io"
)

func main() {
	reader := strings.NewReader(`
	{"Name":"Kayak","Category":"Watersports","Price":279}
	{"Name":"Lifejacket","Category":"Watersports" }
	{"name":"Canoe","category":"Watersports", "price": 100,	"inStock": true }
	`)
	decoder := json.NewDecoder(reader)
	//Запрет неиспользуемых ключей
	//decoder.DisallowUnknownFields()
	for {
		var val Product
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v",
			val.Name, val.Category, val.Price)
		}
	}
}