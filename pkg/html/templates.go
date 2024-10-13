package html

import (
	"embed"
	"html/template"
	"io"
)

//go:embed templates/*.html
var content embed.FS
var templates = template.Must(template.ParseFS(content, "templates/*.html"))

func Render(w io.Writer, name string, data any) {
	templates.ExecuteTemplate(w, name, data)
}
