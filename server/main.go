package main

import (
	"log"
	"mime"
	"net/http"
)

func main() {
	templateCache, err := newTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		templateCache: templateCache,
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
