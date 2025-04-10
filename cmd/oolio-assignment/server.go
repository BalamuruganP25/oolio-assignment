package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initWebServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	http.ListenAndServe(":8080", r)
}
