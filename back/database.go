package back

//Functions for using database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Checks if there is a database file, if not creates one.
func CheckIfDatabaseExists() bool {
	_, errOpening := os.Stat("./books.db")
	if errOpening != nil {
		log.Printf("No database file books.db in ConnecToDb. Creating one.")
		if CreateDb() {
			CreateTableForDb()
			return true
		} else {
			return false
		}
	} else {
		log.Printf("Database file exists.")
		return true
	}
}

// Create database file if there isn't one.
func CreateDb() bool {
	file, errCreating := os.Create("books.db")
	if errCreating != nil {
		log.Printf("Error creating database file books.db in ConnectToDb. %s", errCreating)
		return false
	} else {
		log.Printf("Succesfully created database file books.db")
		file.Close()
		return true
	}
}

// CreateTableForDb, Creates table for database and inserts some test data.
func CreateTableForDb() bool {
	// Connect to database
	database, errDB := ConnectToDb()
	if !errDB {
		log.Printf("Error with connection to database in database.go->PopulateDb")
		return false
	}
	// Create table
	statement, errCreating := database.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, description TEXT)")
	if errCreating != nil {
		log.Printf("Could not create table for books in PopulateDb. %s", errCreating)
		database.Close()
		return false
	}
	_, errExecuting := statement.Exec()
	if errExecuting != nil {
		log.Printf("Error with creating table in PopulateDb. %s", errExecuting)
		database.Close()
		return false
	}
	log.Printf("Succesfully created table in books.db .")
	// Add test data for books and switch to turn it off
	testData := true
	if testData {
		book := Book{
			ID:          0,
			Title:       "The Adventures of Sherlock Holmes",
			Author:      "Doyle, Arthur Conan",
			Description: "A collection of twelve short stories by Arthur Conan Doyle, first published on 14 October 1892",
		}
		AddToDatabase(database, book)
		book2 := Book{
			ID:          0,
			Title:       "The Hollow",
			Author:      "Christie, Agatha",
			Description: `The novel is an example of a "country house mystery"... `,
		}
		AddToDatabase(database, book2)
		book3 := Book{
			ID:          0,
			Title:       "The Big Sleep",
			Author:      "Chandler, Raymond",
			Description: `The title is a euphemism for death; the final pages of the book refer to a rumination about "sleeping the big sleep".`,
		}
		AddToDatabase(database, book3)
	}
	defer database.Close()
	return true

}

// Add book to database using Book struct, return true if successfull.
func AddToDatabase(database *sql.DB, book Book) bool {

	// Adding book data
	statement, _ := database.Prepare("INSERT INTO books ( title, author, description) VALUES (?, ?, ?)")
	_, errExec := statement.Exec(book.Title, book.Author, book.Description)
	if errExec != nil {
		log.Printf("Error with adding new book to database in AddToDatabase. %s.", errExec)
		return false
	}
	log.Printf("Added book %s by %s", book.Title, book.Author)
	return true
}

// Get Books from the Database, return Books in struct array, return true for success.
func GetAllBooksFromDatabase(database *sql.DB) ([]Book, bool) {
	// Variable array to return books with
	var returnBooks []Book
	// Query database
	row, err := database.Query("SELECT * FROM books ORDER BY author")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {

		//Temporary Book
		var tempBook Book
		errDB := row.Scan(&tempBook.ID, &tempBook.Title, &tempBook.Author, &tempBook.Description)
		if errDB != nil {
			log.Printf("Problem querying database in GetAllBooksFromDatabase %s", errDB)
			return nil, false
		}
		// Add temporary book to returning books array.
		returnBooks = append(returnBooks, tempBook)

	}
	// Return books
	return returnBooks, true
}

// Delete from database with id number, returns true for success.
func DeleteBookFromDatabase(database *sql.DB, idToDelete int) bool {
	statement, _ := database.Prepare("DELETE FROM books WHERE id = ?")
	_, errExec := statement.Exec(idToDelete)
	if errExec != nil {
		log.Printf("Error with removing book from database in database.go->DeleteBookFromDatabase. %s.", errExec)
		return false
	}
	//log.Printf("Removed book %s by %s", book.Title, book.Author)
	//return true
	return true
}

// Updating book in database, returns true with success.
func UpdateBookInDatabase(database *sql.DB, bookToUpdate Book) bool {

	statement, _ := database.Prepare("UPDATE books SET title = ?, author = ?, description = ? WHERE id = ?")
	result, errExec := statement.Exec(bookToUpdate.Title, bookToUpdate.Author, bookToUpdate.Description, bookToUpdate.ID)
	if errExec != nil {
		log.Printf("Error updating book in database in database.go->UpdateBookInDatabase. %s.", errExec)
		return false
	}
	// Test if anything changed, 1 = good, anything else is a problem.
	resNumber, _ := result.RowsAffected()
	if resNumber == 1 {
		return true
	} else {
		return false
	}
}

// Connect to database and return pointer to it and true for success.
func ConnectToDb() (*sql.DB, bool) {
	database, errOpening := sql.Open("sqlite3", "./books.db")
	if errOpening != nil {
		log.Printf("Error opening the database file in database.go->ConnectToDb. %s.", errOpening)
		return nil, false
	}

	return database, true
}
