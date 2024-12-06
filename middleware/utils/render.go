package middleware

import (
	"html/template"
	"log"
	"net/http"
)

const _BASE_DIR = "templates/"

func RenderHtml(w http.ResponseWriter, file_name string) {
	tmpl, err := template.ParseFiles(_BASE_DIR + file_name + ".html")
	if err != nil {
		log.Println(err)
		log.Println("Can't render html")
		w.WriteHeader(500)
		RenderHtml(w, "error/500")
		return
	}
	tmpl.Execute(w, nil)
}
