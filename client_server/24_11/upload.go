package main

import (
	"io"
	"fmt"
	"net/http"
)

/*
Поля и метод FileHeader
Name - Это поле возвращает string, содержащую имя файла.
Size - Это поле возвращает значение int64, содержащее размер файла.
Header - Это поле возвращает строку map[string][]string, которая содержит заголовки для части MIME,
	содержащей файл.
Open() - Этот метод возвращает File, который можно использовать для чтения содержимого, связанного с
	заголовком, как показано в следующем разделе.

Поля формы
Value - Это поле возвращает строку map[string][]string, содержащую значения формы.
File - Это поле возвращает map[string][]*FileHeader, содержащий файлы.

*/

func HandleMultipartForm(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Name: %v, City: %v\n", 
		request.FormValue("name"), request.FormValue("city"))
	fmt.Fprintln(writer, "------")

	file, header, err := request.FormFile("files")

	if (err == nil) {
		defer file.Close()
		fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)

		for k, v := range header.Header {
			fmt.Fprintf(writer, "Key: %v, Value: %v\n", k, v)
		}

		fmt.Fprintln(writer, "------")
		io.Copy(writer, file)
	} else {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func HandleMultipartFormSeveralFiles(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10000000)
	fmt.Fprintf(writer, "Name: %v, City: %v\n",
	request.MultipartForm.Value["name"][0],
	request.MultipartForm.Value["city"][0])
	fmt.Fprintln(writer, "------")

	for _, header := range request.MultipartForm.File["files"] {
		fmt.Fprintf(writer, "Name: %v, Size: %v\n", 
			header.Filename, header.Size)

		file, err := header.Open()

		if err == nil {
			defer file.Close()
			fmt.Fprintln(writer, "------")
			io.Copy(writer, file)
		} else {
			http.Error(writer, err.Error(),
			http.StatusInternalServerError)
		return
		}
	}
}
	

func init() {
	// http.HandleFunc("/forms/upload", HandleMultipartForm)
	http.HandleFunc("/forms/upload", HandleMultipartFormSeveralFiles)
}
	