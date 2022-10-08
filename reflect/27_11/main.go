package main

import (
	"reflect"
	// "strings"
	"fmt"
)

type Product struct {
	Name, Category string
	Price          float64
}
var intPtrType = reflect.TypeOf((*int)(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func selectValue(data interface{}, index int) (result interface{}) {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Slice {
		result = dataVal.Index(index).Interface()
	}

	return
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		elemType := reflect.TypeOf(elem)

		if elemType == intPtrType {
			Printfln("Pointer to int: %v", elemValue.Elem().Int())
		} else if elemType == byteSliceType {
			Printfln("Byte slice: %v", elemValue.Bytes())
		} else {
			switch elemValue.Kind() {
			case reflect.Bool:
				var val bool = elemValue.Bool()
				Printfln("Bool: %v", val)
			case reflect.Int:
				var val int64 = elemValue.Int()
				Printfln("Int: %v", val)
			case reflect.Float32, reflect.Float64:
				var val float64 = elemValue.Float()
				Printfln("Float: %v", val)
			case reflect.String:
				var val string = elemValue.String()
				Printfln("String: %v", val)
			}
		}
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	number := 100
	slice := []byte("Alice")
	names := []string{"Alice", "Bob", "Charlie"}
	printDetails(true, 10, 23.30, "Alice", &number, product, slice)

	val := selectValue(names, 1)
	//val := selectValue(names, 1).(string)
	Printfln("Selected: %v, %T", val, val)
}