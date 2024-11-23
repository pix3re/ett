package main

import (
	"net/http"
)

// base routs

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "index.html", *data)
}

// expense routs

func (app *application) newExpense(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "expense-new.html", *data)
}

func (app *application) expenseAdd(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "expense-new.html", *data)
}

// categories routs

func (app *application) newCategory(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "category-new.html", *data)
}

func (app *application) getCategories(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "index.html", *data)
}
