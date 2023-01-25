package  parser

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/cmd/internal/entities"
)

func LoadPage(title string) (*entities.Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	
	if err != nil {
		return nil, err
	}

	return &entities.Page{Title: title, Body: body}, nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, notes interface{}) {
	t, _ := template.ParseFiles(tmpl + ".html")
	err:=t.Execute(w, notes)

	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
