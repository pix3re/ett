package main

import (
	"net/http"
)

func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../static"))

	mux.HandleFunc("GET /{$}", app.home)

	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/images/favicon.ico")
	})

	return mux
}
