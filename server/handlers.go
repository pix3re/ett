package main

import (
	"ett/internal/models"
	"net/http"
)

// base routs

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "index.html", *data)
}

// expense routs

func (app *application) expenseNew(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "expense-new.html", *data)
}

func (app *application) expenseAdd(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "expense-new.html", *data)
}

// categories routs

func (app *application) categoriesNew(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	description := r.PostForm.Get("description")

	c := models.Category{
		Name:        name,
		Description: description,
	}

	id, err := app.category.Insert(c.Name, c.Description)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sendJSON(w, id)
}

func (app *application) categoriesForm(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}

	app.render(w, http.StatusOK, "categories-new.html", *data)
}

func (app *application) categoriesGet(w http.ResponseWriter, r *http.Request) {
	cat, err := app.category.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sendJSON(w, cat)
}
