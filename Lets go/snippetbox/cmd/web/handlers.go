package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	// "path/filepath"
	"strconv"
	
	"snippetbox/pkg/models"

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
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

	data := &templateData{Snippet: record}
	
	// files := []string{
	// 	"./ui/html/show.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
	app.render(w, r, "show.page.tmpl", data)
}

func (app *application)MainHandler(w http.ResponseWriter, r *http.Request) {
	// if the path is not the root path, then return a 404
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	// panic("oops! something went wrong")  // Deliberate panic

	latest, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippets: latest}

	// // Initialize a slice containing the paths to the two files. Note that the
    // // home.page.tmpl file must be the *first* file in the slice.
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

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "internal server error", http.StatusInternalServerError)
	// 	app.serverError(w, err)
	// }
	app.render(w, r, "home.page.tmpl", data)


}