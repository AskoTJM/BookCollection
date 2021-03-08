package back

// Functions for endpoints

import (
	"log"
	"net/http"
)

// Return index.html
func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index")
	// Return index.html
	http.ServeFile(w, r, "./back/frontend/index.html")

}

// GetAllBooks, return all the books in database
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetAllBooks")
}

// UpdateBook, update data on book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateBook")
}

// NewBook, add a book to database
func NewBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("NewBook")
}

// DeleteBook, delete book from database
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteBook")
}
