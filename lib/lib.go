package lib


import (
	"github.com/ali2210/bookish-fiesta/serialization"
	"reflect"
	"context"
	"log"
)

// create book object
var book_row []*serialization.Book

func PublishBook(num int64) []*serialization.Book{
	
	return make([]*serialization.Book, num)
}

type LibraryServiceRPC interface {

	// add book
	
	AddBook(ctx context.Context, books *serialization.Book) (*serialization.Books)
	// search book
	
	SearchBook(ctx context.Context, books *serialization.Book)(*serialization.Books)
	// display book
	
	DisplayBook(ctx context.Context, request *serialization.EmptyRequest) (serialization.StackedBook)
	// delete book
	
	Delete(ctx context.Context, books *serialization.Book) ([]*serialization.Books)
	
	// set stack
	
	SetStackedBook(ctx context.Context, book *serialization.Book) (*serialization.StackedBook)
}

type LibraryServerRPC struct {}


func NewServerRPC() *LibraryServerRPC {return &LibraryServerRPC{}}


func (r *LibraryServerRPC) AddBook(ctx context.Context, book *serialization.Book)(*serialization.Books){


	mylib := make([]serialization.Books, 1)
  
  // for adminstrator    
   
   	mylib[0].Operation = serialization.Stacking_Up
  
 // create empty object type book    
  
 	emptyBook := serialization.Book{}

//   user book ibns empty or empty stack 
// case-I : book ibns compare virtual book ibns (ok)
// case-II : book ibns compare virtual book ibns (error) 

	if !reflect.DeepEqual(book.GetBookIsbn(), emptyBook.GetBookIsbn()) {
	
		mylib[0].Publish = book
	
		mylib[0].Result = serialization.Results_Ok
	
	
		book_row = PublishBook(5) 
  	
	
		if !reflect.DeepEqual(book_row[0], book) { 
	
			book_row[0] = book
	
		}
  
		}else{
	
			mylib[0].Publish = &emptyBook
	
			mylib[0].Result = serialization.Results_Err
	
			log.Printf(" book fields must be have content%v", mylib[0].Publish)
	
			log.Printf(" book fields must have status %v", mylib[0].Result)   
	
			log.Printf("Lib:%v", mylib[0].Publish)
  
		}
    
		return &mylib[0]

	}


	func (r *LibraryServerRPC) SearchBook(ctx context.Context, book *serialization.Book) (*[]serialization.Books){
  	
		//    create virtual object
		emptyBook := &serialization.Book{}
  
	
		libRow := make([]serialization.Books, 1)
  
	
		// add Operation
		libRow[0].Operation = serialization.Stacking_On
 	
		for row := range book_row {
	
			// Case-I if book ins and book on stack ibns are sane then ok
			// Case-II Otherwise, error
			if  reflect.DeepEqual(book.BookIsbn, book_row[row].GetBookIsbn()) {
			
				libRow[0].Result = serialization.Results_Ok
		
				libRow[0].Publish = book_row[row]
			
				return &libRow
	
				}else{
		
					libRow[0].Result = serialization.Results_Err 
		
					libRow[0].Publish = emptyBook
			
					log.Printf("Book fields must be have content:%v\n", libRow[0].Result)
			
					log.Printf("Book:%v searched:%v\n", book.GetBookIsbn(), book_row[row].GetBookIsbn())
			
					}  
  		
				}
  		
				return &libRow
		
			}


		
			func (r *LibraryServerRPC) Delete(ctx context.Context, book *serialization.Book) (*[]serialization.Books){
 
	
		
				libRow := make([]serialization.Books, 1)
 		
				libRow[0].Operation = serialization.Stacking_Down
 			
				for i := range book_row{
	 		
					if reflect.DeepEqual(book_row[i].GetBookIsbn() , book.GetBookIsbn()) {
			
						// remove book from stack
						if status := remove(book_row[i]); status{
					
							libRow[0].Publish = book_row[i]
									
							libRow[0].Result = serialization.Results_Ok
					
							book_row[i] = &serialization.Book{}
			
							}else{
					
								libRow[0].Result = serialization.Results_Err
				
								libRow[0].Publish = book
			
							}
	  			
						}
			
					}
			
					return &libRow
		
				}


			
				func (r *LibraryServerRPC) DisplayBook(ctx context.Context, request *serialization.EmptyRequest) (serialization.StackedBook){

				
					lib := make([]serialization.StackedBook, 1)
	
				
					lib = r.SetStackedBook(ctx, book_row[0])	
	
				
					// sync mutx warning..
	
				
					return lib[0]  

			
				}


			
				func remove(book *serialization.Book)bool{

			
					removeBook := &serialization.Book{}

			
					book = removeBook 

			
					if !reflect.DeepEqual(book, removeBook){

			
						log.Printf(" Book not deleted: %v", book)

			
						return false		
		
					}
		
					return true

				}


		
				func (r *LibraryServerRPC) SetStackedBook(ctx context.Context, book *serialization.Book) ([]serialization.StackedBook){

		
					stacked := make([]serialization.StackedBook, len(book_row)) 

		
					stacked[0].Stackedover = book	

		
					return stacked
	
				}