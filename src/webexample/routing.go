package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "you 've requested the book: %s on Page %s\n", title, page)
	})

	//r.HandleFunc("/books/{title}/page/{page}", ReadBook).Methods("GET").Schemes("http")

	////////////////////////////////////////////////////////////////////////////

	//Methods
	//r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	//r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	//r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	//r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	//Hostnames & Subdomains
	//r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	//Schemes
	//r.HandleFunc("/secure", SecureHandler).Schemes("https")
	//r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	//Path Prefixes & Subrouters
	//bookrouter := r.PathPrefix("/books").Subrouter()
	//bookrouter.HandleFunc("/", AllBooks)
	//bookrouter.HandleFunc("/{title}", GetBook)

	http.ListenAndServe(":80", r)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "you 've requested the book: %s on Page %s\n", title, page)
}
