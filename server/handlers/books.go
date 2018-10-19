package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/MathieuDoyon/bookshelf/server/interfaces"
	"github.com/MathieuDoyon/bookshelf/server/model"
	"github.com/MathieuDoyon/bookshelf/server/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// BooksResource Book router ressources routes
type BooksResource struct {
	Repo interfaces.IBookRepository
	// Repo *repositories.BookRepo
}

// Routes creates a REST router for the books resource
func (rs *BooksResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List)    // GET /books - read a list of books
	r.Post("/", rs.Create) // POST /books - create a new todo and persist it

	r.Route("/{id}", func(r chi.Router) {
		r.Use(rs.BookCtx)        // lets have a books map, and lets actually load/manipulate
		r.Get("/", rs.Get)       // GET /books/{id} - read a single todo by :id
		r.Put("/", rs.Update)    // PUT /books/{id} - update a single todo by :id
		r.Delete("/", rs.Delete) // DELETE /books/{id} - delete a single todo by :id
	})

	return r
}

// List Get all books and filter by query string and sorts
func (rs *BooksResource) List(w http.ResponseWriter, r *http.Request) {
	numberOfPages, err := strconv.ParseInt(r.URL.Query().Get("number_of_pages"), 10, 64)
	yearOfPublication, err := strconv.ParseInt(r.URL.Query().Get("publication_year"), 10, 64)
	rating, err := strconv.ParseInt(r.URL.Query().Get("rating"), 10, 64)

	filters := &model.BookFilter{}
	if r.URL.Query().Get("_id") != "" {
		objectID, err := objectid.FromHex(r.URL.Query().Get("_id"))
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		filters.ID = objectID
	}
	if r.URL.Query().Get("author") != "" {
		filters.Author = r.URL.Query().Get("author")
	}
	if r.URL.Query().Get("genre") != "" {
		filters.Genre = r.URL.Query().Get("genre")
	}
	if r.URL.Query().Get("number_of_pages") != "" {
		filters.NumberOfPages = numberOfPages
	}
	if r.URL.Query().Get("publication_year") != "" {
		filters.YearOfPublication = yearOfPublication
	}
	if r.URL.Query().Get("rating") != "" {
		filters.Rating = rating
	}

	sorting := &model.Sorting{
		Sort:      "publication_year",
		Direction: -1,
	}
	if r.URL.Query().Get("sort") != "" {
		if err := utils.ParseString(r.URL.Query().Get("sort"), &sorting.Sort); err != nil {
			log.Printf("Error parsing sort: %s", err)
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}
	if r.URL.Query().Get("direction") != "" {
		if err := utils.ParseInt32(r.URL.Query().Get("direction"), &sorting.Direction); err != nil {
			log.Printf("Error parsing direction: %s", err)
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}

	books, err := rs.Repo.List(filters, sorting)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, books)
}

// Create Add a new book into collection
func (rs *BooksResource) Create(w http.ResponseWriter, r *http.Request) {
	data := &model.BookRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if created, err := rs.Repo.Create(data.Book); err != nil {
		render.JSON(w, r, err)
	} else {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, created)
	}
}

// Get get a book by ID
func (rs *BooksResource) Get(w http.ResponseWriter, r *http.Request) {
	book := r.Context().Value("book").(*model.Book)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &book)
}

// Update update a book
func (rs *BooksResource) Update(w http.ResponseWriter, r *http.Request) {
	book := r.Context().Value("book").(*model.Book)

	data := &model.BookRequest{Book: book}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	book = data.Book

	if updated, err := rs.Repo.Update(book); err != nil {
		render.JSON(w, r, err)
	} else {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, updated)
	}
}

// Delete remove a book into your bookshelf by ID
func (rs *BooksResource) Delete(w http.ResponseWriter, r *http.Request) {
	book := r.Context().Value("book").(*model.Book)

	if deleted, err := rs.Repo.Delete(book.ID); err != nil {
		render.JSON(w, r, err)
	} else {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, deleted)
	}
}

// BookCtx build book context and inject `model.Book` into request
func (rs *BooksResource) BookCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var book *model.Book
		var err error

		if ID := chi.URLParam(r, "id"); ID != "" {
			book, err = rs.Repo.Get(ID)
		} else {
			render.Render(w, r, ErrNotFound())
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), "book", book)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
