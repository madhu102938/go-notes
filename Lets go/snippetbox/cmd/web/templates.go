package main

import (
	"html/template"
	"path/filepath"
	"time"

	"snippetbox/pkg/models"
	"snippetbox/pkg/forms"
)

type templateData struct {
	Snippet		*models.Snippet
	Snippets    []*models.Snippet
	CurrentYear int
	Form 		*forms.Form
}

func humanDate(t time.Time) string {
	return t.Format(time.DateTime)
}

 // Initialize a template.FuncMap object and store it in a global variable. This is
 // essentially a string-keyed map which acts as a lookup between the names of our
 // custom template functions and the functions themselves.
var functions = template.FuncMap{"humanDate" : humanDate}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Get the file name (last element) from the full file path
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you
        // call the ParseFiles() method. This means we have to use template.New() to
        // create an empty template set, use the Funcs() method to register the
        // template.FuncMap, and then parse the file as normal.
		// New, Funcs return template.Template, thus we can chain them :)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Parse all layout templates in the specified directory
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Parse all partial templates in the specified directory
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// Store the parsed template in the cache with the file name as the key
		cache[name] = ts
	}
	
	// Return the cache map containing all parsed templates
	return cache, nil
}