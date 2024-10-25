package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"path/filepath"
	"snippetbox/pkg/models"
	"strconv"
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

	title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
    expires := "7"

	lastId, err := app.snippets.Insert(title, content, expires)
	
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	redirectURL := fmt.Sprintf("/snippet?id=%d", lastId)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
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

	record, err := app.snippets.Get(id)
	
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	
	fmt.Fprintln(w, record)
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
	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "internal server error", http.StatusInternalServerError)
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "internal server error", http.StatusInternalServerError)
	// 	app.serverError(w, err)
	// }

	latest, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, latest_iter := range latest {
		fmt.Fprintf(w, "%v\n", latest_iter)
	}
}

func (app *application)DownloadHandler(w http.ResponseWriter, r *http.Request) {
    filePath := filepath.Join(".", "ui", "static", "img", "logo.png")
	http.ServeFile(w, r, filePath)
}