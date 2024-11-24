package main

import (
	"net/http"
)

func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../static"))

	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /expense-new", app.expenseNew)
	mux.HandleFunc("POST /expense-add", app.expenseAdd)

	mux.HandleFunc("GET /categories", app.categoriesGet) // return category json to populate lookup ? Or render html in backkend with all values ?
	mux.HandleFunc("GET /categories-new", app.categoriesForm)
	mux.HandleFunc("POST /categories-new", app.categoriesNew)

	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/images/favicon.ico")
	})

	return mux
}
