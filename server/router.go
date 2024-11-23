package main

import (
	"net/http"
)

func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../static"))

	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /new-expense", app.newExpense)
	mux.HandleFunc("POST /expense-add", app.expenseAdd)

	mux.HandleFunc("GET /categories", app.getCategories) // return category json to populate lookup ? Or render html in backkend with all values ?
	mux.HandleFunc("GET /new-category", app.newCategory)

	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/images/favicon.ico")
	})

	return mux
}
