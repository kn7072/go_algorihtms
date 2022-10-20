package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func HTTPSRedirect(writer http.ResponseWriter, request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5500" + request.URL.Path

	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}

	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message",
		http.StatusTemporaryRedirect))

	/*
	Функция FileServer создает обработчик, который будет обслуживать файлы, а
	каталог указывается с помощью функции Dir. (Можно обслуживать файлы
	напрямую, но требуется осторожность, поскольку легко разрешить запросы на выбор
	файлов за пределами целевой папки. Самый безопасный вариант — использовать
	функцию Dir, как показано в этом примере.)

	*/
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))  //  /files/store.html

	// https://getacert.com/  для генерации сертификата
	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer",
			"certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()

	// err := http.ListenAndServe(":5000", nil)
	err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))

	if err != nil {
		Printfln("Error: %v", err.Error())
	}
	// https://localhost:5500/templates/products.html
	// https://localhost:5500/templates/edit.html?index=2
	// https://localhost:5500/files/upload.html
	// https://localhost:5500/cookies
}
