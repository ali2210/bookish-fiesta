package lib


import (
	"github.com/ali2210/bookish-fiesta/serialization"
	"reflect"
	"context"
	"log"
)

type option serialization.Results
var book_row []*serialization.Book


func PublishBook() []*serialization.Book{
	book_row = make([]*serialization.Book, 5)
	return book_row
}


type LibraryServiceRPC interface {
	AddBook(ctx context.Context, books *serialization.Book) (*serialization.Books)
	SearchBook(ctx context.Context, books *serialization.Book)(*serialization.Books)
	DisplayBook(ctx context.Context, books *serialization.Book) (*serialization.StackedBook)
	Delete(ctx context.Context, books *serialization.Book) (*serialization.Books)
}

type LibraryServerRPC struct {}

func NewServerRPC() *LibraryServerRPC {return &LibraryServerRPC{}}

func (r *LibraryServerRPC) AddBook(ctx context.Context, book *serialization.Book)(*serialization.Books){

  mylib := make([]serialization.Books, 5)
  emptyBook := serialization.Book{}
  if reflect.DeepEqual(book, emptyBook) {
	mylib[0].Publish = &emptyBook
	mylib[0].Result = serialization.Results_Err
	log.Printf(" book fields must be have content", mylib[0]) 
	return &mylib[0]
  }
  mylib[0].Publish = book
  mylib[0].Result = serialization.Results_Ok
  for row := range book_row{
	  if reflect.DeepEqual(book_row[row], book) && reflect.DeepEqual(book_row[row],emptyBook) {
		  book_row[row] = book
	  }
  }
  return &mylib[0]
}

func (r *LibraryServerRPC) SearchBook(ctx context.Context, book *serialization.Book) (*serialization.Books){
  emptyBook := serialization.Book{}
  libRow := make([]serialization.Books, 5)
  if  reflect.DeepEqual(book, emptyBook) {
	libRow[0].Result = serialization.Results_Err 
	log.Printf(" book fields must be have content:", libRow[0].Publish)
	return &libRow[0]
  }
  for row := range book_row {
	if reflect.DeepEqual(book.BookIsbn, book_row[row].BookIsbn){
		libRow[0].Result = serialization.Results_Ok 
		log.Println("match:", book)
		return &libRow[0]
	}else{
		libRow[0].Result = serialization.Results_Err 
		log.Printf(" this is not available on the database:", book)
	}
  }
  return &libRow[0]
}

func (r *LibraryServerRPC) Delete(ctx context.Context, book *serialization.Book) (serialization.Books){}

func (r *LibraryServerRPC) DisplayBook(ctx context.Context, book *serialization.Book) (serialization.StackedBook){}

