# BookCollection
Book Collection, Buutti Backend Developer Task
March 2021 

IMPORTANT!
Frontend doesn't work with Firefox. Tested on Windows 10 and Ubuntu 20.04 Virtual Machine.
Some issue with fetch.

Tested working:
Chrome works fine on both Windows 10 and VM Ubuntu 20.04.
Brave and Edge in Windows 10.
Opera in Ubuntu 20.04

Backend: 
Programmed in Go 1.15.6
Using Windows 10 Home 2004 and VSCode 

Run with 'go run main.go' or with the executable.
If the browser doesn't automatically open go to address 127.0.0.1:8080



If there is no books.db, it will be created with three books as sample data.

Endpoints:

GET /
Serves index.html for browser

GET /books
Responses with JSON of all the books in database

POST /books
Add new book to database with JSON

example: 
    {
        "title": "New Title",
        "author": "New Author",
        "description": "New Description"
    }

PATCH /books/{index}
Update books data in database with JSON using it's database index number

Notice! 
Currently not sending key/value pair clears that value in database.
Should ignore changes and only clear previous value with '"key":""' 

Also endpoint address {index} overrides possible "id":"{index}" from JSON.

DELETE /books/{index}
Removes book from database

Frontend:

Made in HTML/Javascript

