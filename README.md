# BookCollection
Book Collection, Buutti Backend Developer Task
March 2021 

Backend: 
Programmed in Go 1.15.6
With VSCode

Run it with 'go run main.go' or with the executable.
If the browser doesn't automatically open go to address localhost:8080

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
Should ignore changes and only clear with '"key":""' 

Also endpoint address {index} overrides possible "id":"{index}" from JSON.

DELETE /books/{index}
Removes book from database

Frontend:

Made in HTML/Javascript

