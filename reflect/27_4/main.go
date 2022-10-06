package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type Product struct {
	Name, Category string
	Price          float64
}

type Customer struct {
	Name, City string
}

type Payment struct {
	Currency string
	Amount   float64
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)

		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldValue := elemValue.Field(i)
				fieldDetails = append(fieldDetails,
					fmt.Sprintf("%v: %v", fieldName, fieldValue))
			}

			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

func printDetailsOLD(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
		case Product:
			Printfln("Product: Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		case Customer:
			Printfln("Customer: Name: %v, City: %v",
				val.Name, val.City)
		}
	}
}

func main() {
	product := Product{Name: "Kayak",
		Category: "Watersports",
		Price:    279,
	}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment{Currency: "USD", Amount: 100.50}

	printDetailsOLD(product, customer)
	printDetails(product, customer, payment)
}
