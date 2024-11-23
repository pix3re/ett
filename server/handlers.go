package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "index.html", *data)
}
