package utilities

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, filename string, data any, statusCode int) {
	tmpl, err := template.ParseFiles(filename)
	
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	buf.WriteTo(w)
}