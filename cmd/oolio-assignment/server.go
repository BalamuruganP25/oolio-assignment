package main

import (
	"net/http"
	"oolio-assignment/pkg/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initWebServer(processConfig handler.ProcessConfig) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	// Your other routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":8089", r)
}
