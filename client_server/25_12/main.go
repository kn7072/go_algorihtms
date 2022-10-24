package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

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

/*
Таблица 25-5 Клиентские поля и методы

Transport - Это поле используется для выбора транспорта, который будет
	использоваться для отправки HTTP-запроса. Пакет net/http
	предоставляет транспорт по умолчанию.
CheckRedirect - Это поле используется для указания пользовательской политики
	для обработки повторяющихся перенаправлений, как описано в
	разделе «Управление перенаправлениями».
Jar - Это поле возвращает файл CookieJar, который используется для
	управления файлами cookie, как описано в разделе «Работа с файлами cookie».
Timeout - Это поле используется для установки тайм-аута для запроса,
	указанного как time.Duration.
Do(request) - Этот метод отправляет указанный Request, возвращая Response и
	error, указывающую на проблемы с отправкой запроса.
CloseIdleConnections() - Этот метод закрывает все бездействующие HTTP-запросы,
	которые в настоящее время открыты и не используются.
Get(url) - Этот метод вызывается функцией Get, описанной в таблице 25-3.
Head(url) - Этот метод вызывается функцией Head, описанной в таблице 25-3.
Post(url, contentType, reader) - Этот метод вызывается функцией Post, описанной в таблице 25-3.
PostForm(url, data) - Этот метод вызывается функцией PostForm, описанной в таблице 25-3.

*/

/*
Таблица 25-6 Полезные поля и методы запроса

Method - Это строковое поле указывает метод HTTP, который будет использоваться
	для запроса. Пакет net/http определяет константы для методов HTTP,
	таких как MethodGet и MethodPost.
URL - В этом поле URL указывается URL-адрес, на который будет отправлен
	запрос. Структура URL определена в главе 24.
Header - Это поле используется для указания заголовков запроса. Заголовки
	указываются в map[string][]string, и поле будет nil, когда значение
	Request создается с использованием синтаксиса литеральной структуры.
ContentLength - Это поле используется для установки заголовка Content-Length с
	использованием значения int64.
TransferEncoding - Это поле используется для установки заголовка Transfer-Encoding с
	использованием среза строк.
Body - Это поле ReadCloser указывает источник тела запроса. Если у вас есть
	Reader, который не определяет метод Close, то можно использовать
	функцию io.NopCloser для создания ReadCloser, метод Close которого
	ничего не делает.

*/

/*
Таблица 25-7 Функция для анализа значений URL

Parse(string) - Этот метод анализирует string в URL. Результатами являются значение URL
	и error, указывающая на проблемы с разбором string.
*/

/*
Таблица 25-8 Удобные функции net/http для создания запросов

NewRequest(method, url, reader) - Эта функция создает новый Reader, настроенный с
	указанным методом, URL-адресом и телом. Функция также возвращает ошибку, указывающую на проблемы с
	созданием значения, включая синтаксический анализ URL-адреса, выраженного в виде строки.
NewRequestWithContext(context, method, url, reader) - Эта функция создает новый Reader, который будет
	отправлен в указанном контексте. Контексты описаны в главе 30.
*/

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])

	if err == nil {
		response, err := http.Post("http://localhost:5000/echo",
			"application/json",
			strings.NewReader(builder.String()))

		if err == nil && response.StatusCode == http.StatusOK {
			io.Copy(os.Stdout, response.Body)
			defer response.Body.Close()
		} else {
			Printfln("Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

	if err == nil {
		reqURL, err := url.Parse("http://localhost:5000/echo")
		if err == nil {
			req := http.Request{
				Method: http.MethodPost,
				URL:    reqURL,
				Header: map[string][]string{
					"Content-Type": {"application.json"},
				},
				Body: io.NopCloser(strings.NewReader(builder.String())),
			}
			response, err := http.DefaultClient.Do(&req)

			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Parse Error: %v", err.Error())
		}
	} else {
		Printfln("Encoder Error: %v", err.Error())
	}

	if err == nil {
		req, err := http.NewRequest(http.MethodPost,
			"http://localhost:5000/echo",
			io.NopCloser(strings.NewReader(builder.String())))

		if err == nil {
			req.Header["Content-Type"] = []string{"application/json"}
			response, err := http.DefaultClient.Do(req)

			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Request Init Error: %v", err.Error())
		}
	} else {
		Printfln("Encoder Error: %v", err.Error())
	}

	//------------------------------------------
	/*
		Таблица 25-10 Функция конструктора Cookie Jar в пакете net/http/cookiejar
			New(options) - Эта функция создает новый CookieJar, настроенный с помощью структуры
				Options, описанной далее. Функция также возвращает error, сообщающую о
				проблемах с созданием jar.

		Таблица 25-9 Методы, определяемые интерфейсом CookieJar
			SetCookies(url, cookies) - Этот метод сохраняет срез *Cookie для указанного URL-адреса.
			Cookies(url) - Этот метод возвращает срез *Cookie, содержащий файлы cookie, которые
				должны быть включены в запрос для указанного URL-адреса.
	*/
	jar, err := cookiejar.New(nil)
	if err == nil {
		http.DefaultClient.Jar = jar
	}

	for i := 0; i < 3; i++ {
		req, err := http.NewRequest(http.MethodGet,
			"http://localhost:5000/cookie", nil)

		if err == nil {
			response, err := http.DefaultClient.Do(req)
			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Request Init Error: %v", err.Error())
		}
	}
}
