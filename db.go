package main

import (
	"database/sql"
	)

type Book struct{
	Name string
	Year string
	Length string
}

func dbConnect() error {
	var err error
	db, err = sql.Open("mysql", "root:lena1997@/lab?charset=utf8")
	if err != nil {
		return err
	}
	return nil
}

func dbAddBook(name, year, length string) error {
	stmt, err := db.Prepare("INSERT books SET book_name=?,book_year=?,book_length=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, year, length)
	if err != nil {
		return err
	}
	return nil
}

func dbGetBooks() ([]Book, error) {
	var books []Book
	stmt, err := db.Prepare("SELECT book_name, book_year, book_length FROM books")
	if err != nil {
		return books, err
	}

	res, err := stmt.Query()
	if err != nil {
		return books, err
	}

	var tempBook Book
	for res.Next() {
		err = res.Scan(&tempBook.Name, &tempBook.Year, &tempBook.Length)
		if err != nil {
			return books, err
		}
		books = append(books, tempBook)
	}

	return books, err
}