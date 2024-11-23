package main

import (
	"ett/internal/models"
	"log"
	"mime"
	"net/http"
)

func main() {
	db, err := openDB()

	if err != nil {
		log.Fatal(err)
	}

	err = initDB(db, "../SQL/schema.sql")

	if err != nil {
		log.Fatal(err)
	}

	templateCache, err := newTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		templateCache: templateCache,
		expense:       &models.ExpenseModel{DB: db},
		category:      &models.CategoryModel{DB: db},
	}

	log.Println("Server running on http://localhost:8000")

	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")

	server := http.Server{
		Addr:    ":8000",
		Handler: app.router(),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
