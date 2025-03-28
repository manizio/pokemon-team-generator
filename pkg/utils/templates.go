package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates, _ = template.ParseGlob("templates/*.html")
}

func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	err := templates.ExecuteTemplate(w, template, data)

	if err != nil {
		fmt.Printf(">> erro: %v", err.Error())
	}
}
