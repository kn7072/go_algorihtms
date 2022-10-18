package main

import (
	"html/template"
	"net/http"
	"strconv"
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

type Context struct {
	Request *http.Request
	Data []Product
}

var htmlTemplates *template.Template

func HandleTemplateRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	if path == "" {
		path = "products.html"
	}

	t := htmlTemplates.Lookup(path)

	if t == nil {
		http.NotFound(writer, request)
	} else {
		err := t.Execute(writer, Context{request, Products})

		if (err != nil) {
			http.Error(writer, err.Error(),	http.StatusInternalServerError)
		}
	}
}

func init() {
	var err error

	htmlTemplates = template.New("all")

	htmlTemplates.Funcs(map[string]interface{} {
		"intVal": strconv.Atoi,
	})

	htmlTemplates, err = htmlTemplates.ParseGlob("templates/*.html")
	if err == nil {
		http.Handle("/templates/", http.StripPrefix("/templates/",
			http.HandlerFunc(HandleTemplateRequest)))
	} else {
		panic(err)
	}
}
	