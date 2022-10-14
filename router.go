package main

import (
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() chi.Router {
	router := chi.NewRouter()
	middleware.DefaultLogger = middleware.RequestLogger(
		&middleware.DefaultLogFormatter{
			Logger: newLogger(),
		},
	)
	router.Use(middleware.Logger)
	router.HandleFunc("/*", Index)
	router.Post("/update", postDb)
	router.Get("/index", getData)
	return router

}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "chi-log:", log.Lshortfile)
}
