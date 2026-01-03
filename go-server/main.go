package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "Form post successfull\n")
	name := r.FormValue("name")
	pass := r.FormValue("pass")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Password: %s\n", pass)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello to the go server!!!\n")
}

func main() {
	staticFile := http.FileServer(http.Dir("static"))
	http.Handle("/", staticFile)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("starting server at port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
