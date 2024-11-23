package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
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

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "../DB/test.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDB(db *sql.DB, schemaPath string) error {
	schema, err := os.ReadFile(schemaPath)

	if err != nil {
		return err
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return err
	}

	return nil
}
