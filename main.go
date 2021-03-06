package main

/* Buutti Backend Developer Task
// 6.3.2021
*/

import (
	"fmt"
	browser "github.com/pkg/browser"
	"log"
	"net/http"

	sw "github.com/AskoTJM/BookCollection/back"
)

func main() {

	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

	// Open Browser when run
	const url = "localhost"
	browser.OpenURL(url)

	fmt.Println("Test1")
}
