package tmpl

import "html/template"

func Funcs() template.FuncMap {
	return template.FuncMap{
		"hi": Hi,
	}
}

// --

func Hi() string {
	return "hi"
}
