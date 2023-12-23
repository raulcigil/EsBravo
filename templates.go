package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

var templatesPath string = "templates/"

// Define the template registry struct
type Template struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base", data)
}

func RegisterTemplates(e *echo.Echo) {
	templates := make(map[string]*template.Template)
	templates["index"] = template.Must(template.ParseFiles(templatesPath+"base.html", templatesPath+"index.html"))
	//templates["about.html"] = template.Must(template.ParseFiles("templates/about.html", "templates/base.html"))
	e.Renderer = &Template{
		templates: templates,
	}
}
