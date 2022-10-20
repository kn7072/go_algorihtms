package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Product struct {
	Name, Category string
	Price          float64
}

var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

/*
методы для HTTP-запросов
	Get(url) - Эта функция отправляет запрос GET на указанный URL-адрес HTTP или
		HTTPS. Результатом являются ответ и error, сообщающая о проблемах с запросом.
	Head(url) - Эта функция отправляет запрос HEAD на указанный URL-адрес HTTP или
		HTTPS. Запрос HEAD возвращает заголовки, которые были бы возвращены
		для запроса GET. Результатом являются Response и error, сообщающая о проблемах с запросом.
	Post(url, contentType, reader) - Эта функция отправляет запрос POST на указанный URL-адрес HTTP или
		HTTPS с указанным значением заголовка Content-Type. Содержимое формы предоставляется указанным Reader.
		Результатом являются Response и error, сообщающая о проблемах с запросом.
	PostForm(url, data) - Эта функция отправляет запрос POST на указанный URL-адрес HTTP или
		HTTPS с заголовком Content-Type, установленным в application/x-www-form-urlencoded. 
		Содержимое формы предоставляется с помощью map[string][]string. Результатом являются Response 
		и error, сообщающая о проблемах с запросом.
*/

func main() {
	Printfln("Starting HTTP Server")
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	response, err := http.Get("http://localhost:5000/html")
	if err == nil {
		response.Write(os.Stdout)
	} else {
		Printfln("Error: %v", err.Error())
	}

	// http://localhost:5000/echo
	// http://localhost:5000/html
	// http://localhost:5000/json
}
