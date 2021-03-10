package back

// Functions for endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Return index.html
func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index")
	// Return index.html
	http.ServeFile(w, r, "./back/frontend/index.html")

}

// GetAllBooks, return all the books in database
func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	//log.Printf("GetAllBooks")
	// Set Header Content-Type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	database, errConn := ConnectToDb()
	defer database.Close()
	if !errConn {
		log.Printf("Error connecting to database in functions.go->GetAllBooks.")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	}

	// Get All Books from database
	returnedBooks, errDB := GetAllBooksFromDatabase(database)
	if !errDB {
		log.Printf("Error with receiving books in functions.go->GetAllBooks.")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	} else {
		anon, _ := json.Marshal(returnedBooks)
		n := len(anon)
		response := string(anon[:n])
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", response)
		//log.Print(returnedBooks)
	}
}

// UpdateBook, update data on book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateBook")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	//setHeaders(&w)
	var tempBook Book

	// Decoding JSON->Book Struct
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	errDecode := dec.Decode(&tempBook)
	if errDecode != nil {
		log.Printf("Error decoding json in functions.go->NewBook() %s", errDecode)
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Get Book ID
	vars := mux.Vars(r)
	sesId := vars["book"]
	// Turn Id to int
	bookId, errConv := strconv.Atoi(sesId)
	if errConv != nil {
		log.Printf("Problem with turning")
	}
	// Replacing possible ID from JSON with one of the endpoints.
	tempBook.ID = bookId

	// Connect to Database
	database, errConn := ConnectToDb()
	// Close database connection when not needed anymore
	defer database.Close()
	if !errConn {
		log.Printf("Connection to database failed in functions.go->UpdateBook()")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	}

	// Update the book in database
	errDB := UpdateBookInDatabase(database, tempBook)
	if !errDB {
		log.Printf("Connection to database failed in functions.go->UpdateBook()")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Book updated successfully.")
	}
}

// NewBook, add a book to database
func NewBook(w http.ResponseWriter, r *http.Request) {
	//log.Printf("NewBook()")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var tempBook Book

	// Decoding JSON->Book Struct
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	errDecode := dec.Decode(&tempBook)
	if errDecode != nil {
		log.Printf("Error decoding json in functions.go->NewBook() %s", errDecode)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")

	}

	// Connect to Database
	database, errConn := ConnectToDb()
	// Close database connection when not needed anymore
	defer database.Close()
	if !errConn {
		log.Printf("Connection to database failed in functions.go->NewBook()")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	}

	// Add the book to database
	errDB := AddToDatabase(database, tempBook)
	if !errDB {
		log.Printf("Adding to database failed in functions.go->NewBook()")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Book added successfully.")

}

// DeleteBook, delete book from database
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteBook")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get Book ID
	vars := mux.Vars(r)
	sesId := vars["book"]
	// Turn Id to int
	bookId, errConv := strconv.Atoi(sesId)
	if errConv != nil {
		log.Printf("Problem with turning")
	}
	// Connect To Database
	database, errConn := ConnectToDb()
	log.Printf("Connected to database")
	// Close database connection when not needed anymore
	defer database.Close()
	if !errConn {
		log.Printf("Connection to database failed in functions.go->DeleteBook()")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	}
	// Delete book
	log.Printf("Ready to delete book.")
	DeleteBookFromDatabase(database, bookId)
	log.Printf("Returning from deleting book.")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Book removed successfully.")

}

// Sets header values
func setHeaders(w *http.ResponseWriter) {
	// Big no no for real use, but CORS is being tricky and this is just local
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Request-Method", "*")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

}
