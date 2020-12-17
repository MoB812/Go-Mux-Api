package main

// Import Packages
import (
	"log"
	"net/http"

	"./bookslib"

	"github.com/gorilla/mux"
)

// Main Function
func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Mock Data
	bookslib.Books = append(bookslib.Books, bookslib.Book{ID: "1", ISBN: "573896", Title: "Book One", Author: &bookslib.Author{FirstName: "John", LastName: "Doe"}})
	bookslib.Books = append(bookslib.Books, bookslib.Book{ID: "2", ISBN: "123456", Title: "Book Two", Author: &bookslib.Author{FirstName: "Billy", LastName: "Joe"}})

	// Route Endpoints
	router.HandleFunc("/api/books", bookslib.GetBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", bookslib.GetBook).Methods("GET")
	router.HandleFunc("/api/books", bookslib.CreateBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", bookslib.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", bookslib.DeleteBook).Methods("DELETE")

	// Listen on Port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}
