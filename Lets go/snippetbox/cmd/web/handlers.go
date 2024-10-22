package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"path/filepath"
)

func (app *application)CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a new snippet"))
}

func (app *application)SnippetHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Display a snippet"))
	// receiving arguements
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// fmt.Fprintln(w, "not a valid id")
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Displaying a specific snippet with ID %v\n", id)
	// writing to response writer with fprintf
}

func (app *application)MainHandler(w http.ResponseWriter, r *http.Request) {
	// if the path is not the root path, then return a 404
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w)
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
		// app.errorLog.Println(err.Error())
		// http.Error(w, "internal server error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "internal server error", http.StatusInternalServerError)
		app.serverError(w, err)
	}
}

func (app *application)DownloadHandler(w http.ResponseWriter, r *http.Request) {
    filePath := filepath.Join(".", "ui", "static", "img", "logo.png")
	http.ServeFile(w, r, filePath)
}