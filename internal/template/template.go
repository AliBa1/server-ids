package template

import (
	"io"
	"text/template"
)

type Templates struct {
	templates         *template.Template
	LastRenderedBlock string
	LastRenderedData  interface{}
}

type ReturnData struct {
	Error string
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("internal/views/*.html")),
	}
}

func NewTestTemplate() *Templates {
	return &Templates{
		templates: template.New("test"),
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	t.LastRenderedBlock = name
	t.LastRenderedData = data
	return t.templates.ExecuteTemplate(w, name, data)
}
