package main

import (
	"fmt"
	"net/http"
)

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "This is the hello page!")
}

func FormHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm error: %v", err)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name: %s\naddress: %s\n", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	//Route for the home page
	http.Handle("/", fileServer)

	// Route for the hello page
	http.HandleFunc("/hello", HelloHandle)

	//Route for the form page
	http.HandleFunc("/form", FormHandle)

	fmt.Println("Starting the server on the port 8000...")
	err := http.ListenAndServe(":8000", nil)
	checkErrNil(err)
}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
