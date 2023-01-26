package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/redirect", RedirectHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Println(w, "POST request succesfull")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	fmt.Fprintln(w, "Name :", name)
	fmt.Fprintln(w, "Surname :", surname)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, "https://www.artstation.com/artwork/ZGd80m", http.StatusPermanentRedirect)
}
