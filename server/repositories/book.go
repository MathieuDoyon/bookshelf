package repositories

// func Book{}
import (
	"context"
	"log"

	"github.com/MathieuDoyon/bookshelf/server/db"
	"github.com/MathieuDoyon/bookshelf/server/interfaces"
	"github.com/MathieuDoyon/bookshelf/server/model"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
)

// BookRepo book repository
type BookRepo struct {
	interfaces.IBookRepository
}

// List query database to return list of books
func (repo *BookRepo) List(filters *model.BookFilter, sorting *model.Sorting) ([]model.Book, error) {
	collection := db.Database.Collection("books")

	filterDoc := bson.NewDocument()
	if filters.Author != "" {
		filterDoc.Append(bson.EC.String("author", filters.Author))
	}
	if filters.Genre != "" {
		filterDoc.Append(bson.EC.String("genre", filters.Genre))
	}
	if filters.NumberOfPages != 0 {
		filterDoc.Append(bson.EC.Int64("number_of_pages", filters.NumberOfPages))
	}
	if filters.YearOfPublication != 0 {
		filterDoc.Append(bson.EC.Int64("publication_year", filters.YearOfPublication))
	}
	if filters.Rating != 0 {
		filterDoc.Append(bson.EC.Int64("rating", filters.Rating))
	}

	var opts []findopt.Find

	sort := "_id"
	if sorting.Sort != "" {
		sort = sorting.Sort
	}

	var direction int32 = -1
	if sorting.Direction != 0 {
		direction = sorting.Direction
	}

	opts = append(opts, findopt.Sort(bson.NewDocument(bson.EC.Int32(sort, direction))))
	opts = append(opts, findopt.Limit(10))

	cur, err := collection.Find(
		nil,
		filterDoc,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var books []model.Book

	for cur.Next(nil) {
		book := model.Book{}
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal("Decode error ", err)
		}
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("Cursor error ", err)
	}

	return books, nil
}

// Create add a new book into database
func (repo *BookRepo) Create(book *model.Book) (*model.Book, error) {
	collection := db.Database.Collection("books")

	book.ID = objectid.New()
	_, err := collection.InsertOne(nil, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// Get get a book by ID into database
func (repo *BookRepo) Get(ID string) (*model.Book, error) {
	collection := db.Database.Collection("books")

	objectID, err := objectid.FromHex(ID)
	if err != nil {
		return nil, err
	}
	var book *model.Book

	idDoc := bson.NewDocument(bson.EC.ObjectID("_id", objectID))

	err = collection.FindOne(nil, idDoc).Decode(&book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// Update update a book by ID
func (repo *BookRepo) Update(book *model.Book) (*model.Book, error) {
	collection := db.Database.Collection("books")

	idDoc := bson.NewDocument(bson.EC.ObjectID("_id", book.ID))

	_, err := collection.UpdateOne(
		nil,
		idDoc,
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("author", book.Author),
				bson.EC.String("genre", book.Genre),
				bson.EC.Int64("number_of_pages", book.NumberOfPages),
				bson.EC.Int64("publication_year", book.YearOfPublication),
				bson.EC.Int64("rating", book.Rating),
			),
			bson.EC.SubDocumentFromElements("$currentDate",
				bson.EC.Boolean("modifiedAt", true),
			),
		),
	)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// Delete remove a book from database
func (repo *BookRepo) Delete(ID objectid.ObjectID) (int64, error) {
	collection := db.Database.Collection("books")

	idDoc := bson.NewDocument(bson.EC.ObjectID("_id", ID))

	res, err := collection.DeleteOne(nil, idDoc)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
