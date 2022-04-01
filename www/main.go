package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
 
	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {

}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_article", create)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()

}
