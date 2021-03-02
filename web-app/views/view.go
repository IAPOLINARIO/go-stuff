package views

import "html/template"

// NewView creates a new view object
func NewView(layout string, files ...string) *View {
	files = append(files, "views/layout/mainpage.gohtml")

	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View is the struct that holds the template definition
type View struct {
	Template *template.Template
	Layout   string
}
