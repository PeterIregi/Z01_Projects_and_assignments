package utils

import (
	"net/http"
	"html/template"
)

type ErrorData struct {
	Code  int
	Message string
}

func ErrorHandler(w http.ResponseWriter,r *http.Request, code int, msg string){
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/error_page.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := ErrorData{
		Code: code,
		Message: msg,
	}

	tmpl.Execute(w, data)
}
