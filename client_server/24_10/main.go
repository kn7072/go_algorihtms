package main

// Использование обработчика удобной маршрутизации

import (
	"fmt"
	"io"
	"net/http"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func main() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message",
		http.StatusTemporaryRedirect))

	/*
	Ключом к этой функции является использование nil в качестве аргумента
	функции ListenAndServe, например:
	Это включает обработчик по умолчанию, который направляет запросы
	обработчикам на основе правил, установленных с помощью функций, описанных в таблице 24-9.
	
	Handle(pattern, handler) - Эта функция создает правило, которое вызывает указанный метод ServeHTTP указанного
		Handler для запросов, соответствующих шаблону.
	HandleFunc(pattern, handlerFunc) -  Эта функция создает правило, которое вызывает указанную функцию для запросов,
 		соответствующих шаблону. Функция вызывается с аргументами ResponseWriter и Request.

	Таблица 24-10 The net/http Functions for Creating Request Handlers
		FileServer(root) - Эта функция создает Handler, который выдает ответы с помощью функции
			ServeFile. См. раздел Создание статического HTTP-сервера для примера, который обслуживает файлы.
		NotFoundHandler() - Эта функция создает Handler, который выдает ответы с помощью функции NotFound.
		RedirectHandler(url, code) - Эта функция создает Handler, который выдает ответы с помощью функции Redirect.
		StripPrefix(prefix, handler) - Эта функция создает Handler, который удаляет указанный префикс из URL-адреса
			запроса и передает запрос указанному Handler. Подробнее см. в разделе «Создание	статического HTTP-сервера».
		TimeoutHandler(handler, duration, message) - Эта функция передает запрос указанному Handler, но генерирует ответ об ошибке,
			если ответ не был получен в течение указанного времени.
	*/
		err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
