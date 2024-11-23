package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		fmt.Printf("Template page %s was not found\n", page)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "layout", data)
	if err != nil {
		fmt.Printf("Failed to load template - %s\nErr - %s", page, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
