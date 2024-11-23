package main

import "html/template"

type application struct {
	templateCache map[string]*template.Template
}

type templateData struct {
}
