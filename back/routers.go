package back

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var routes = Routes{
	// Endpoint to serve index.html
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	// Endpoint to serve all the books
	Route{
		"GetAllBooks",
		"GET",
		"/books",
		GetAllBooks,
	},
	// Endpoint to update a book
	Route{
		"UpdateBook",
		"PATCH",
		"/books/{book}",
		UpdateBook,
	},
	// Endpoint to add new book
	Route{
		"AddBook",
		"POST",
		"/books",
		NewBook,
	},
	// Endpoint to delete book
	Route{
		"DeleteBook",
		"DELETE",
		"/books/{book}",
		DeleteBook,
	},
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
