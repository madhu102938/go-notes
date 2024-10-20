// webprogrammming, sending http requests

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// you can access the documentation with `go doc FunctionName`

var templates *template.Template
var homeTemplate *template.Template
// using a global variable for templates, so that it can be accessed by all handlers

func main() {
	PopularTemplates()

	http.HandleFunc("/profile", ProfileHandler)
	http.HandleFunc("/profile/", ProfileHandler)
	// if you don't add '/' at the end, then it will not match the url

	http.HandleFunc("/", HomeHandler)
	// if nothing is matched, then '/' is used, thus it needs to at bottom

	fmt.Println("server started on 8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
	// nil is used because we are using default server mux
}

// PopularTemplates parses all the templates in the directory.
func PopularTemplates() {
	var err error
	templates, err = template.ParseGlob("D:/web_develop/Simon_game/*.html")
	// goes to the directory and parses all the html files

	if err != nil {
		log.Fatal("templates couldn't be read", err)
	}

	homeTemplate = templates.Lookup("index.html")
	// this is used to get the template from the templates
}

// ProfileHandler handles "/profile" page.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "me developer, me go lang, plz hire me :)")
}

// HomeHandler handles "/" page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "hehe starting page, about time")
	if r.Method == "GET" {
		homeTemplate.Execute(w, nil)
		// applies the template to the response writer
	}
}
