package main

import (
	"ett/internal/models"
	"html/template"
)

type application struct {
	templateCache map[string]*template.Template
	expense       *models.ExpenseModel
	category      *models.CategoryModel
}

type templateData struct {
}
