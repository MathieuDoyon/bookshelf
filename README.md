#Bookshelf

This is an example of an API using go-chi. You can use docker-compose to start server with hot reload. [Realize](https://gorealize.io/) is install inside docker to server the API with live reload

Dependencies:
 - [Chi (Router)](https://github.com/go-chi/chi)
 - [Testify (Test & Mock framework)](https://github.com/stretchr/testify)
 - [Mockery (Mock generator)](https://github.com/vektra/mockery)
 - [Mongo DB driver](github.com/mongodb/mongo-go-driver)
 - [Viper](github.com/spf13/viper)
 - [Realize](https://gorealize.io/)

 Get Started:

 - [Install](#install)
 - [Starting Server](#start)
 - [Using](#Using)
 - [TODOs](#TODOs)

 ----------
[Install](#install)
-------

Clone the source into your `$GOPATH/src/github.com/MathieuDoyon/bookshelf`
```bash
git clone git@github.com:MathieuDoyon/bookshelf.git
```

Setup dependencies (recommended way to install is using [dep](https://github.com/golang/dep))
```bash
make install
# or use dep
dep ensure
```
 ----------
[Starting Server](#start)
-------
It will start the server into docker with live reload
```
make serve
```

If you want to run the server inside your terminal instead of running it into docker, you need have a running instance of mongo and environment var exported.
```
# Export all environment config to terminal
export $(cat ./.env | xargs)
```
 ----------
[Using](#Using)
-------

HTTPie
```
# Add a new book into bookshelf
http POST :8080/books < ./fixtures/book.json

http POST :8080/books author="Mathieu Doyon" genre=Fiction number_of_pages:=345 publication_year:=2020 rating:=5

# Get list of books
http GET :8080/books/ 

# Get list of book with filters
# All book properties can be added as query string to filter the request.
# author, genre, number_of_pages, publication_year, rating
http GET :8080/books/ rating==4 sort==author direction==-1

# Get a specific book by Mongo Object ID
http GET :8080/books/{ID}

# Update a book
http PUT :8080/books/{ID} genre="SCI FI & FANTASY"

```

 ----------
[Test](#Test)
-------
Run test with makefile
```
make test
```

 ----------
[TODOs](#TODOs)
-------
- [ ] Dockerfile (production)
