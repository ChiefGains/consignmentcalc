package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Index!")
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/home.html")
}

func main() {

	fmt.Println("Now serving on port 8080")

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("web/static")))

	http.Handle("/static/", static)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/home", handleHome)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
