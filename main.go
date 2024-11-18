package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, err := template.ParseFiles(lp, fp)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/images/favicon.ico")
	})

	mux.HandleFunc("/", serveTemplate)

	log.Println("Server running on https://localhost:8000")

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}
