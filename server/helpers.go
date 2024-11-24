package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

func (app *application) sendJSON(w http.ResponseWriter, data any) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonStr)
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

func (app *application) serverError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	fmt.Println(err)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
