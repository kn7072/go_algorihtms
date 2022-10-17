package main

// Создание прослушивателя и обработчика HTTP

import (
	"fmt"
	"io"
	"net/http"
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

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Method: %v", request.Method)
	Printfln("URL: %v", request.URL)
	Printfln("HTTP Version: %v", request.Proto)
	Printfln("Host: %v", request.Host)

	for name, val := range request.Header {
		Printfln("Header: %v, Value: %v", name, val)
	}

	Printfln("---")

	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/test":
		writer.WriteHeader(http.StatusNotFound)
	case "/message":
		io.WriteString(writer, sh.message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
}

/*
Удобные функции net/http
	ListenAndServe(addr, handler) - Эта функция начинает прослушивать HTTP-запросы по указанному адресу и
	передает запросы указанному обработчику.
	ListenAndServeTLS(addr, cert, key, handler) - Эта функция начинает прослушивать HTTPS-запросы. Аргументы - это адрес

Метод, определяемый интерфейсом обработчика
	ServeHTTP(writer, request) - Этот метод вызывается для обработки HTTP-запроса. Запрос описывается значением
	Request, а ответ записывается с использованием ResponseWriter, оба из которых принимаются в качестве параметров.

Метод ResponseWriter
	Header() - Этот метод возвращает Header, который является псевдонимом для map[string][]string,
		который можно использовать для установки заголовков ответа.
	WriteHeader(code) - Этот метод устанавливает код состояния для ответа, заданного как int. Пакет net/http
		определяет константы для большинства кодов состояния.
	Write(data) - Этот метод записывает данные в тело ответа и реализует интерфейс Writer.

Основные поля, определяемые структурой запроса - request
	Method - В этом поле указывается метод HTTP (GET, POST и т. д.) в виде строки. Пакет net/http определяет
	константы для методов HTTP, таких как MethodGet и MethodPost.
	URL - Это поле возвращает запрошенный URL-адрес, выраженный в виде URL значения.
	Proto - Это поле возвращает string, указывающую версию HTTP, используемую для запроса.
	Host - Это поле возвращает string, содержащую запрошенный хост.
	Header - Это поле возвращает значение Header, которое является псевдонимом для map[string][]string и
	содержит заголовки запроса. Ключи карты — это имена заголовков, а значения — срезы строк,
	содержащие значения заголовков.
	Trailer - Это поле возвращает строку map[string], содержащую любые дополнительные заголовки,
	включенные в запрос после тела.
	Body - Это поле возвращает ReadCloser, представляющий собой интерфейс, сочетающий метод Read
	интерфейса Reader с методом Close интерфейса Closer, оба из которых описаны в главе 22.

Полезные поля и методы, определяемые структурой URL
	Scheme - Это поле возвращает компонент схемы URL.
	Host - Это поле возвращает хост-компонент URL-адреса, который может включать порт.
	RawQuery - Это поле возвращает строку запроса из URL-адреса. Используйте метод Query для преобразования
	строки запроса в карту.
	Path - Это поле возвращает компонент пути URL-адреса.
	Fragment - Это поле возвращает компонент фрагмента URL без символа #.
	Hostname() - Этот метод возвращает компонент имени хоста URL-адреса в виде string.
	Port() - Этот метод возвращает компонент порта URL-адреса в виде string.
	Query() - Этот метод возвращает строку map[string][]string (карту со строковыми ключами и строковыми
		значениями срезов), содержащую поля строки запроса.
	User() - Этот метод возвращает информацию о пользователе, связанную с запросом, как описано в главе 30.
	String() - Этот метод возвращает string представление URL-адреса.

Удобные функции ответа
	Error(writer, message, code) - Эта функция устанавливает для заголовка указанный код, устанавливает для заголовка
		Content-Type значение text/plain и записывает сообщение об ошибке в ответ. Заголовок
		X-Content-Type-Options также настроен так, чтобы браузеры не могли интерпретировать
		ответ как что-либо, кроме текста.
	NotFound(writer, request) - Эта функция вызывает Error и указывает код ошибки 404.
	Redirect(writer, request, url, code) - Эта функция отправляет ответ о перенаправлении на указанный URL-адрес и с указанным
		кодом состояния.
	ServeFile(writer, request, fileName) - Эта функция отправляет ответ, содержащий содержимое указанного файла. Заголовок
		Content-Type устанавливается на основе имени файла, но его можно переопределить,
		явно установив заголовок перед вызовом функции. См. раздел «Создание статического
		HTTP-сервера» для примера, который обслуживает файлы.

*/

func main() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello,	World"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
