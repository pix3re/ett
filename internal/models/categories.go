package models

import (
	"database/sql"
)

type Category struct {
	ID          int
	Name        string
	Description string
}

type CategoryModel struct {
	DB *sql.DB
}
