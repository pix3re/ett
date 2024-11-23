package main

import (
	"fmt"
	"html/template"
	"path/filepath"
)

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../templates/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Printf("Parsed files: %s\n", name)

		ts, err := template.New(name).ParseFiles("../templates/layout.html")
		if err != nil {
			return nil, err
		}

		//ts, err = ts.ParseGlob("../templates/partials/*.html")
		//if err != nil {
		//	return nil, err
		//}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
