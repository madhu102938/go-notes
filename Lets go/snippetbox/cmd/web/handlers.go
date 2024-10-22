package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"path/filepath"
)

var count int = 0
func (app *application)CreateHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		count++;
	}()

	if r.Method != http.MethodPost {
		if count == 0 {
			w.Header().Set("allow-hehe", http.MethodPost)
			// it will automatically change to `Allow-Hehe`
		}
		if count == 1 {
			w.Header().Add("Allow", "hehe")
		}
		// fmt.Println(w.Header().Get("Allow"))
		if count > 2 {
			w.Header().Del("Allow")
		}
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a new snippet"))
}

func (app *application)SnippetHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Display a snippet"))
	// receiving arguements
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		fmt.Fprintln(w, "not a valid id")
		return
	}

	fmt.Fprintf(w, "Displaying a specific snippet with ID %v\n", id)
	// writing to response writer with fprintf
}

func (app *application)MainHandler(w http.ResponseWriter, r *http.Request) {
	// if the path is not the root path, then return a 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// // Initialize a slice containing the paths to the two files. Note that the
    // home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (app *application)DownloadHandler(w http.ResponseWriter, r *http.Request) {
    filePath := filepath.Join(".", "ui", "static", "img", "logo.png")
	http.ServeFile(w, r, filePath)
}