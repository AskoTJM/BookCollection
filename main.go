package main

/* Buutti Backend Developer Task
// 6.3.2021
*/

import (
	"log"
	"net/http"

	sw "github.com/AskoTJM/bookcollection/back"
	"github.com/pkg/browser"
)

func main() {

	if sw.CheckIfDatabaseExists() {
		// Open Browser with when run
		const url = "http://127.0.0.1:8080/"
		err := browser.OpenURL(url)
		// Show error if something goes wrong.
		if err != nil {
			log.Printf("%s", err)
		}
	}

	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
