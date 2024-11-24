package models

import (
	"database/sql"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryModel struct {
	DB *sql.DB
}

func (m *CategoryModel) GetAll() ([]*Category, error) {
	query := `SELECT id, name, description FROM categories`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []*Category

	for rows.Next() {
		c := &Category{}

		err = rows.Scan(&c.ID, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (m *CategoryModel) Insert(pName, pDescription string) (int, error) {
	query := `INSERT INTO categories (name, description) VALUES (?, ?)`

	res, err := m.DB.Exec(query, pName, pDescription)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
