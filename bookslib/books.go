package bookslib

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Books - Init books variable as a slice Book struct
var Books []Book

// GetBooks Function
func GetBooks(res http.ResponseWriter, req *http.Request) {
	// Set the header to app/json
	res.Header().Set("Content-Type", "application/json")
	// Encode the books struct into JSON format
	json.NewEncoder(res).Encode(Books)
}

// GetBook Function
func GetBook(res http.ResponseWriter, req *http.Request) {
	// Set the header to app/json
	res.Header().Set("Content-Type", "application/json")

	// Get parameters from incoming request
	params := mux.Vars(req)

	// Loop through books and find with id
	for _, item := range Books {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	// ID not found, return empty book struct
	json.NewEncoder(res).Encode(&Book{})
}

// CreateBook Function
func CreateBook(res http.ResponseWriter, req *http.Request) {
	// Set the header to app/json
	res.Header().Set("Content-Type", "application/json")

	// Create new book variable
	var book Book
	// Decode book params passed into api
	_ = json.NewDecoder(req.Body).Decode(&book)
	// Generate ID
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock id for testing, can generate same id
	// Append to global books struct inventory
	Books = append(Books, book)

	// Send response of book created
	json.NewEncoder(res).Encode(book)
}

// UpdateBook Function
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	// Set the header to app/json
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			// Create new book variable
			var book Book
			// Decode book params passed into api
			_ = json.NewDecoder(req.Body).Decode(&book)
			// Generate ID
			book.ID = params["id"]
			// Append to global books struct inventory
			Books = append(Books, book)

			// Send response of book created
			json.NewEncoder(res).Encode(book)
			return
		}
	}
	// Send response of book created
	json.NewEncoder(res).Encode(Books)
}

// DeleteBook Function
func DeleteBook(res http.ResponseWriter, req *http.Request) {
	// Set the header to app/json
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	// Send response of book created
	json.NewEncoder(res).Encode(Books)
}
