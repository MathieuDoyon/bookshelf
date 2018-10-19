package interfaces

import (
	"net/http"

	"github.com/MathieuDoyon/bookshelf/server/model"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type IBookRepository interface {
	List(filters *model.BookFilter, sorting *model.Sorting) ([]model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Get(ID string) (*model.Book, error)
	Update(book *model.Book) (*model.Book, error)
	Delete(ID objectid.ObjectID) (int64, error)
	BookCtx(next http.Handler) http.Handler
}
