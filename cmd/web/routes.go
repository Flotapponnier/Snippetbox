package main

import (
	"net/http"
)

func (app *application) routes(cfg config) *http.ServeMux {

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux

}
