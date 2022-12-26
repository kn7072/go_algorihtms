package main

import (
	"encoding/json"
	"fmt"
	"io"
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

/*
Таблица 25-4 Поля и методы, определяемые структурой Response
	StatusCode - Это поле возвращает код состояния ответа, выраженный как int.
	Status - Это поле возвращает string, содержащую описание статуса.
	Proto - Это поле возвращает string, содержащую ответный HTTP-протокол.
	Header - Это поле возвращает строку map[string][]string, содержащую заголовки ответа.
	Body - Это поле возвращает ReadCloser, который является Reader,
		определяющим метод Close и обеспечивающим доступ к телу ответа.
	Trailer - Это поле возвращает строку map[string][]string, содержащую трейлеры ответов.
	ContentLength - Это поле возвращает значение заголовка Content-Length,
		преобразованное в значение int64.
	TransferEncoding - Это поле возвращает набор значений заголовка Transfer-Encoding.
	Close - Это логическое поле возвращает значение true, если ответ содержит
		заголовок Connection, для которого установлено значение close, что
		указывает на то, что HTTP-соединение должно быть закрыто.
	Uncompressed - Это поле возвращает значение true, если сервер отправил сжатый ответ,
		который был распакован пакетом net/http.
	Request - Это поле возвращает Request, который использовался для получения
		ответа. Структура Request описана в главе 24.
	TLS - В этом поле содержится информация о соединении HTTPS.
	Cookies() - Этот метод возвращает []*Cookie, который содержит заголовки Set-
		Cookie в ответе. Структура Cookie описана в главе 24.
	Location - Этот метод возвращает URL-адрес из ответа заголовка Location и error,
		указывающую, что ответ не содержит этот заголовок.
	Write(writer) - Этот метод записывает сводку ответа на указанный Writer.

*/

func main() {
	Printfln("Starting HTTP Server")
	go http.ListenAndServe(":5000", nil)
	time.Sleep(2 * time.Second)

	response, err := http.Get("http://localhost:5000/html")
	if (err == nil && response.StatusCode == http.StatusOK) {

		data, err := io.ReadAll(response.Body)

		if (err == nil) {
			defer response.Body.Close()
			os.Stdout.Write(data)
		}
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(),
			response.StatusCode)
	}

	response, err = http.Get("http://localhost:5000/json")
	if (err == nil && response.StatusCode == http.StatusOK) {
		defer response.Body.Close()

		data := []Product{}
		err = json.NewDecoder(response.Body).Decode(&data)

		if err == nil {
			for _, p := range data {
				Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
			}
		} else {
			Printfln("Decode error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}

	formData := map[string][]string {
		"name": { "Kayak "},
		"category": { "Watersports"},
		"price": { "279"},
	}
		
	response, err = http.PostForm("http://localhost:5000/echo", formData)
	if (err == nil && response.StatusCode == http.StatusOK) {
		defer response.Body.Close()
		io.Copy(os.Stdout, response.Body)
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(),response.StatusCode)
	}

	// http://localhost:5000/echo
	// http://localhost:5000/html
	// http://localhost:5000/json
}
