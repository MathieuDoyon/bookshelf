package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MathieuDoyon/bookshelf/server/db"
	"github.com/MathieuDoyon/bookshelf/server/handlers"
	"github.com/MathieuDoyon/bookshelf/server/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("bookshelf")
	viper.AutomaticEnv()

	db.Configure()
}

func main() {
	defer db.Client.Disconnect(nil)

	bookResource := handlers.BooksResource{
		Repo: &repositories.BookRepo{},
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("When the seagulls follow the trawler, it's because they think sardines will be thrown into the sea."))
	})
	r.Mount("/books", bookResource.Routes())

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("server is listening on port :8080")
}
