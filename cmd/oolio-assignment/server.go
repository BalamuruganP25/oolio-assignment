package main

import (
	"net/http"
	"oolio-assignment/pkg/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initWebServer(processConfig handler.ProcessConfig) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	http.ListenAndServe(":8080", r)
}
