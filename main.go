package main

import (
	"net/http"
		"log"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	)

var db *sql.DB

func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}

		books, err := dbGetBooks()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, books)
	}
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		name := r.Form.Get("name")
		year := r.Form.Get("year")
		length := r.Form.Get("length")
		err := dbAddBook(name, year, length)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addBookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
