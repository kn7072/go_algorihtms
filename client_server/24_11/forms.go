package main

import (
	"net/http"
	"strconv"
)

/*
Поля данных и методы формы запроса
Form - Это поле возвращает строку map[string][]string, содержащую
	проанализированные данные формы и параметры строки запроса. Перед чтением
	этого поля необходимо вызвать метод ParseForm.
PostForm - Это поле похоже на Form, но исключает параметры строки запроса, поэтому в карте
	содержатся только данные из тела запроса. Перед чтением этого поля необходимо
	вызвать метод ParseForm.
MultipartForm - Это поле возвращает составную форму, представленную с помощью структуры
	Form, определенной в пакете mime/multipart. Перед чтением этого поля
	необходимо вызвать метод ParseMultipartForm.
FormValue(key) - Этот метод возвращает первое значение для указанного ключа формы и возвращает
	пустую строку, если значение отсутствует. Источником данных для этого метода
	является поле Form, а вызов метода FormValue автоматически вызывает ParseForm
	или ParseMultipartForm для анализа формы.
PostFormValue(key) - Этот метод возвращает первое значение для указанного ключа формы и возвращает
	пустую строку, если значение отсутствует. Источником данных для этого метода
	является поле PostForm, а вызов метода PostFormValue автоматически вызывает
	ParseForm или ParseMultipartForm для анализа формы.
FormFile(key) - Этот метод обеспечивает доступ к первому файлу с указанным в форме ключом.
	Результатами являются File и FileHeader, оба из которых определены в пакете
	mime/multipart, и error. Вызов этой функции приводит к вызову функций
	ParseForm или ParseMultipartForm для анализа формы.
ParseForm() - Этот метод анализирует форму и заполняет поля Form и PostForm. Результатом
	является error, которая описывает любые проблемы синтаксического анализа.
ParseMultipartForm(max) - Этот метод анализирует составную форму MIME и заполняет поле MultipartForm.
	Аргумент указывает максимальное количество байтов, выделяемых для данных
	формы, а результатом является error, описывающая любые проблемы с обработкой формы.
*/

func ProcessFormData(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		index, _ := strconv.Atoi(request.PostFormValue("index"))
		p := Product{}
		p.Name = request.PostFormValue("name")
		p.Category = request.PostFormValue("category")
		p.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
		Products[index] = p
	}

	http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
}

func init() {
	http.HandleFunc("/forms/edit", ProcessFormData)
	// https://localhost:5500/templates/edit.html?index=2
}
