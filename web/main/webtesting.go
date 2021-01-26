package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"code/github.com/consignmentcalc/internal/generator"
	"code/github.com/consignmentcalc/internal/files"
)

type UserSession struct{
	user *generator.User
	sync.Pool
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func handleInventory(w http.ResponseWriter, r *http.Request) {
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
