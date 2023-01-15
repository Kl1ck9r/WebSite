package services

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

func RenderTemplate(w http.ResponseWriter, tmpl string, p *entities.Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}
