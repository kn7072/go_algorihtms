package main

import "encoding/json"

type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:"offer,string"`
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


type Product struct {
	Name, Category string
	Price float64
}

var Kayak = Product {
	Name: "Kayak",
	Category: "Watersports",
	Price: 279,
}

var Products = []Product {
	{ "Kayak", "Watersports", 279 },
	{ "Lifejacket", "Watersports", 49.95 },
	{ "Soccer Ball", "Soccer", 19.50 },
	{ "Corner Flags", "Soccer", 34.95 },
	{ "Stadium", "Soccer", 79500 },
	{ "Thinking Cap", "Chess", 16 },
	{ "Unsteady Chair", "Chess", 75 },
	{ "Bling-Bling King", "Chess", 1200 },
}
