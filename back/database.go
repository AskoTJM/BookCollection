package back

//Functions for using database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CheckIfDatabaseExists() bool {
	_, errOpening := os.Stat("./books.db")
	if errOpening != nil {
		log.Printf("No database file books.db in ConnecToDb. Creating one.")
		if CreateDb() {
			PopulateDb()
			return true
		} else {
			return false
		}
	} else {
		log.Printf("Database file exists.")
		return true
	}
}

func CreateDb() bool {
	// Create database file if there isn't one.
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

// PopulateDb, to have some test data.
func PopulateDb() bool {

	// Open database
	database, errOpening := sql.Open("sqlite3", "./books.db")
	if errOpening != nil {
		log.Printf("Error with opening the database file in PopulateDb. %s", errOpening)
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

	book := Book{
		ID:          0,
		Title:       "ATitle",
		Author:      "AAuthor",
		Description: "ADesc",
	}
	AddToDatabase(book)
	book2 := Book{
		ID:          0,
		Title:       "BTitle",
		Author:      "BAuthor",
		Description: "BDesc",
	}
	AddToDatabase(book2)
	book3 := Book{
		ID:          0,
		Title:       "CTitle",
		Author:      "CAuthor",
		Description: "CDesc",
	}
	AddToDatabase(book3)
	defer database.Close()
	return true

}

func AddToDatabase(book Book) bool {
	database, errOpening := sql.Open("sqlite3", "./books.db")
	if errOpening != nil {
		log.Printf("Error with opening the database file in AddToDatabase. %s.", errOpening)
		return false
	}
	defer database.Close()
	statement, _ := database.Prepare("INSERT INTO books ( title, author, description) VALUES (?, ?, ?)")
	_, errExec := statement.Exec(book.Title, book.Author, book.Description)
	if errExec != nil {
		log.Printf("Error with adding new book to database in AddToDatabase. %s.", errExec)
		return false
	} else {
		log.Printf("Added book %s by %s", book.Title, book.Author)
		return true
	}
	//defer database.Close()
	//return true
}
