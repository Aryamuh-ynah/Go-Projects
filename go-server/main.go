package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	
	fmt.Fprintf(w, "Hello!")
}

func fromHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "\nName = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}



func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", fromHandler)
	http.HandleFunc("/Hello", helloHandle)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}