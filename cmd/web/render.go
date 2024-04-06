package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template

	// get template from cache
	if app.config.useCache {
		if cachedTemplate, ok := app.templateCache[t]; ok {
			tmpl = cachedTemplate
		}
	}

	// if there was not template in cache, build it
	if tmpl == nil {
		templateFromDisk, err := app.buildTemplate(t)
		if err != nil {
			log.Println("Error building template:", err)
			return
		}

		log.Println("Building the '" + t + "' template from disk...")
		tmpl = templateFromDisk
	}

	// if template data is nil, initialize it
	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) buildTemplate(t string) (*template.Template, error) {
	templates := []string{
		"./templates/base.layout.html",
		"./templates/partials/navbar.partial.html",
		"./templates/partials/header.partial.html",
		"./templates/partials/footer.partial.html",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return nil, err
	}

	app.templateCache[t] = tmpl

	return tmpl, nil
}
