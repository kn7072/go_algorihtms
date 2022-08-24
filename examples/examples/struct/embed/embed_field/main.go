package main

import (
	"fmt"
)

type Embeded struct {
	A, B int
}

func (t Embeded) MethodTestEmbeded() {
	fmt.Println("MethodTestEmbeded", t.A, t.B)
}

func (t *Embeded) MethodTestEmbededPtr(newA int) {
	t.A = newA
	fmt.Println("MethodTestEmbededPtr", t.A, t.B)
}

// func (t Embeded) MethodTestEmbededPtr(newb int) {
// 	t.A = newb
// 	fmt.Println("MethodTestEmbededPtr", t.A, t.B)
// }

type TestEmbedValueReceiver struct {
	Embeded // встроенное поле - значение
}

type TestEmbedPtrReceiver struct {
	*Embeded // встроенное поле - указатель
}

func main() {
	embededValueReceiver := TestEmbedValueReceiver{Embeded: Embeded{A: 1, B: 2}}
	embededValueReceiver.MethodTestEmbeded()
	embededValueReceiver.MethodTestEmbededPtr(5)
	fmt.Println(embededValueReceiver.A)

	embededPtrReceiver := TestEmbedPtrReceiver{Embeded: &Embeded{A: 3, B: 4}}
	embededPtrReceiver.MethodTestEmbeded()
	embededPtrReceiver.MethodTestEmbededPtr(6)
	fmt.Println(embededPtrReceiver.A)
}
