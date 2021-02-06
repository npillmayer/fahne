package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/npillmayer/fahne/webserver/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Get("/users/{userID}", handlers.TestHandler)

	webRoot := os.Getenv("WEBROOT")
	log.Printf("WEBROOT=%s", webRoot)
	filesDir := http.Dir(filepath.Join(webRoot, "static"))
	handlers.FileServer(r, "/static", filesDir)

	log.Println("serving :3333")
	http.ListenAndServe(":3333", r)
}
