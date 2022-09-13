package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// type DiscountedProduct struct {
// 	*Product `json:"product"`
// 	Discount float64 `json:"-"`
// }

//Пропуск неназначенных полей
// стр 695
// type DiscountedProduct struct {
// 	*Product `json:"product,omitempty"`
// 	Discount float64 `json:"-"`
// }

//Принудительное кодирование полей как строк
type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:",string"`
}

func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if (dp.Product != nil) {
		m := map[string]interface{} {
			"product": dp.Name,
			"cost": dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}			

func main(){
	dp := DiscountedProduct{
		Product: &Kayak,
		Discount: 10.50,
	}
	dp2 := DiscountedProduct{ Discount: 10.50 }

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(&dp)
	encoder.Encode(&dp2)

	namedItems := []Named{ &dp, &Person{ PersonName: "Alice"}}
	encoder.Encode(namedItems)

	fmt.Print(writer.String())
}
