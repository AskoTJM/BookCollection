# BookCollection
Book Collection, Buutti Backend Developer Task
March 2021 

Install:
Download repo as Zip.
Extract.

Options for running backend.

1. On Windows: run bookcollection.exe
2. On Linux: 
    chmod +x bookcollection  
    ./bookcollection
3. If GO 1.15.6 is installed on machine
    go run main.go

If the browser doesn't automatically open go to address 127.0.0.1:8080

If there is no books.db, it will be created with three books as sample data.

IMPORTANT!
Frontend doesn't work with Firefox. Tested on Windows 10 and Ubuntu 20.04 Virtual Machine.
For some reason it doesn't like fetch-command.

Tested working:
Chrome on both Windows 10 and Ubuntu 20.04 Virtual Machine.
Brave and Edge in Windows 10.
Opera in Ubuntu 20.04 Virtual Machine

Backend: 

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

