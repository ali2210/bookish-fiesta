package main

import (
	"github.com/ali2210/bookish-fiesta/lib"
	"github.com/ali2210/bookish-fiesta/serialization"
	"context"
	"fmt"
)


//  For playaround add single book  

func main() {
	// create rpc server object 
	server := lib.NewServerRPC()
	
	// add your book
	collection := server.AddBook(context.Background(), &serialization.Book{
		BookID : "book-0", BookTitle : "Category Theory for Programmers", BookIsbn:"9780464243878", BookAuthor : " Bartosz Milewski", BookPublishDate : "2018 09 17"})
	fmt.Println("\n================   Book Add in Collection ============\n", collection)
	fmt.Println("\n =======================================================================")
	//  check libaray stack
	displayCollection := server.DisplayBook(context.Background(), &serialization.EmptyRequest{})
	fmt.Println("\n ================ Book Collection ============\n", displayCollection)
	fmt.Println("\n =======================================================================")
	// // search book 
	searchCollection := server.SearchBook(context.Background(), collection.Publish)
	fmt.Println("\n ================   Search Book  for Book ============\n", searchCollection)
	fmt.Println("\n =======================================================================")
	// // delete book 
	deleteCollection := server.Delete(context.Background(), collection.Publish)
	fmt.Println("\n ================ Delete Book ============\n", deleteCollection)
	fmt.Println("\n =======================================================================")
}