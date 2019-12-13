package controllers

import (
	"html/template"
	"net/http"
)

//HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	type ViewData struct {
		Title string
	}

	data := ViewData{
		Title: "Текст",
	}

	tmpl := template.Must(template.ParseFiles("Views/Layout.html", "Views/Home.html"))
	tmpl.Execute(w, data)
}
