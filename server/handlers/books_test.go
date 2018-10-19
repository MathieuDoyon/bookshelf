package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/MathieuDoyon/bookshelf/server/interfaces/mocks"
	"github.com/MathieuDoyon/bookshelf/server/model"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestBookList(t *testing.T) {
	var expected []model.Book

	expected = append(expected, model.Book{
		Author: "Mathieu Doyon",
	})

	filters := &model.BookFilter{}

	sorting := &model.Sorting{
		Sort:      "publication_year",
		Direction: -1,
	}

	repoMock := &mocks.IBookRepository{}
	repoMock.On("List", filters, sorting).Return(expected, nil) // mock the expectation

	bookResource := BooksResource{
		Repo: repoMock,
	}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:8080/books", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/books", bookResource.List)

	r.ServeHTTP(w, req)

	response := []model.Book{}

	json.NewDecoder(w.Body).Decode(&response)

	// assert that the expectations were met
	assert.Equal(t, expected, response)
	repoMock.AssertNumberOfCalls(t, "List", 1)
	repoMock.AssertExpectations(t)
}

// func TestBookGet(t *testing.T) {
// 	expected := &model.Book{
// 		Author: "Mathieu Doyon",
// 	}

// 	repoMock := &mocks.IBookRepository{}
// 	repoMock.On("Get", "fooooooo").Return(expected, nil) // mock the expectation

// 	bookResource := BooksResource{
// 		Repo: repoMock,
// 	}

// 	// call the code we are testing
// 	req := httptest.NewRequest("GET", "http://localhost:8080/books/fooooooo", nil)
// 	w := httptest.NewRecorder()

// 	r := chi.NewRouter()
// 	// r.HandleFunc("/books/{id}", bookResource.Get)
// 	r.Route("/{id}", func(r chi.Router) {
// 		r.Use(bookResource.BookCtx)
// 		r.Get("/", bookResource.Get)
// 	})

// 	r.ServeHTTP(w, req)

// 	response := model.Book{}

// 	json.NewDecoder(w.Body).Decode(&response)

// 	// assert that the expectations were met
// 	// assert.Equal(t, expected, response)
// 	repoMock.AssertNumberOfCalls(t, "Get", 1)
// 	repoMock.AssertExpectations(t)
// }
