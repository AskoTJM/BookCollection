package main

/* Buutti Backend Developer Task
// 6.3.2021
*/

import (
	"fmt"
	browser "github.com/pkg/browser"
	"log"
	"net/http"
)

func main() {
	// Open Browser when run
	const url = "http://golang.org/"
	browser.OpenURL(url)

	fmt.Println("Test1")
}
