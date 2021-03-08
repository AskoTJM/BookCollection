package back

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

func ConnectToDb() bool {
	database, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Printf("Error with opening the database file books.db .")
		// Create database file if there isn't one.
		file, err2 := os.Create("books.db")
		if err2 != nil {
			log.Printf("Error creating database file books.db .")
			return false
		} else {
			log.Printf("Succesfully created database file books.db")
			file.Close()
			return true
		}
	} else {
		log.Print(database)
		return true
	}
	return true
}

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
