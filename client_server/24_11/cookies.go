package main

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
Функция net/http для настройки файлов cookie
	SetCookie(writer, cookie) -  Эта функция добавляет заголовок Set-Cookie к указанному ResponseWriter. Файл cookie
		описывается с помощью указателя на структуру Cookie, которая описана далее.

Поля, определяемые структурой cookie
	Name - Это поле представляет имя файла cookie, выраженное в виде строки.
	Value - Это поле представляет значение файла cookie, выраженное в виде строки.
	Path - В этом необязательном поле указывается путь к файлу cookie.
	Domain - В этом необязательном поле указывается host/domain, для которого будет установлен файл cookie.
	Expires - В этом поле указывается срок действия файла cookie, выраженный в виде значения time.Time.
	MaxAge - В этом поле указывается количество секунд до истечения срока действия файла cookie, выраженное как int.
	Secure - Когда это bool поле имеет значение true, клиент будет отправлять файл cookie только через
		соединения HTTPS.
	HttpOnly - Когда это bool поле имеет значение true, клиент предотвратит доступ кода JavaScript к файлу cookie.
	SameSite - В этом поле указывается политика перекрестного происхождения для файла cookie с использованием
		констант SameSite, которые определяют SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode и SameSiteNoneMode.

Методы запроса файлов cookie
	Cookie(name) - Этот метод возвращает указатель на значение Cookie с указанным именем и error,
		указывающую на отсутствие соответствующего файла cookie.
	Cookies() - Этот метод возвращает срез указателей Cookie.

*/

func GetAndSetCookie(writer http.ResponseWriter, request *http.Request) {
	counterVal := 1
	counterCookie, err := request.Cookie("counter")

	if err == nil {
		counterVal, _ = strconv.Atoi(counterCookie.Value)
		counterVal++
	}

	http.SetCookie(writer, &http.Cookie{
		Name:  "counter",
		Value: strconv.Itoa(counterVal),
	})

	if len(request.Cookies()) > 0 {
		for _, c := range request.Cookies() {
			fmt.Fprintf(writer, "Cookie Name: %v, Value: %v", c.Name, c.Value)
		}
	} else {
		fmt.Fprintln(writer, "Request contains no cookies")
	}
}

func init() {
	http.HandleFunc("/cookies", GetAndSetCookie)
}
