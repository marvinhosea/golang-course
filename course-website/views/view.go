package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func NewView(layout string, files ...string) *View {
	files = append(files,
		layoutFiles()...
	)
	templ, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}

	return &View{
		Template: templ,
		Layout: layout,
	}
}

type View struct {
	Template *template.Template
	Layout string
}

func (view *View) Render(writer http.ResponseWriter, data interface{}) error {
	return view.Template.ExecuteTemplate(writer, view.Layout, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob("views/layouts/*.gohtml")

	if err != nil {
		panic(err)
	}

	return files
}