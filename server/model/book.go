package model

import (
	"errors"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Book book structure
type Book struct {
	ID                objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Author            string            `bson:"author" json:"author"`
	Genre             string            `bson:"genre" json:"genre"`
	NumberOfPages     int64             `bson:"number_of_pages" json:"number_of_pages"`
	YearOfPublication int64             `bson:"publication_year" json:"publication_year"`
	Rating            int64             `bson:"rating" json:"rating"`
}

type BookFilter struct {
	ID                objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Author            string            `bson:"author" json:"author"`
	Genre             string            `bson:"genre" json:"genre"`
	NumberOfPages     int64             `bson:"number_of_pages" json:"number_of_pages"`
	YearOfPublication int64             `bson:"publication_year" json:"publication_year"`
	Rating            int64             `bson:"rating" json:"rating"`
}

type BookRequest struct {
	*Book

	ProtectedID string `json:"_id"` // override '_id' json to have more control
}

func (b *BookRequest) Bind(r *http.Request) error {
	// b.Book is nil if no Book fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if b.Book == nil {
		return errors.New("missing required Book fields.")
	}

	// b.User is nil if no Userpayload fields are sent in the request. In this app
	// this won't cause a panic, but checks in this Bind method may be required if
	// b.User or futher nested fields like b.User.Name are accessed elsewhere.

	// just a post-process after a decode..
	b.ProtectedID = "" // unset the protected ID
	// b.Book.Title = strings.ToLower(b.Book.Title) // as an example, we down-case
	return nil
}
