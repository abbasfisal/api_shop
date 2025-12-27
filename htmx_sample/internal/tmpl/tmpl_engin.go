package tmpl

import "html/template"

func TmplEngine() *template.Template {
	t := template.New("").Funcs(Funcs())

	must := func(patt string) {
		template.Must(t.ParseGlob(patt))
	}

	must("internal/views/*.html")

	return t
}
