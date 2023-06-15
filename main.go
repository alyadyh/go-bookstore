package main

import (
	"go-bookstore/config"
	"go-bookstore/controllers/bookcontroller"
	"go-bookstore/controllers/genrecontroller"
	"go-bookstore/controllers/homecontroller"
	"go-bookstore/controllers/staffcontroller"
	"go-bookstore/controllers/transactioncontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/books", bookcontroller.Index)
	http.HandleFunc("/books/detail", bookcontroller.Detail)
	http.HandleFunc("/books/add", bookcontroller.Add)
	http.HandleFunc("/books/edit", bookcontroller.Edit)
	http.HandleFunc("/books/delete", bookcontroller.Delete)

	http.HandleFunc("/genres", genrecontroller.Index)
	http.HandleFunc("/genres/add", genrecontroller.Add)
	http.HandleFunc("/genres/edit", genrecontroller.Edit)
	http.HandleFunc("/genres/delete", genrecontroller.Delete)

	http.HandleFunc("/staff", staffcontroller.Index)
	http.HandleFunc("/staff/add", staffcontroller.Add)
	http.HandleFunc("/staff/edit", staffcontroller.Edit)
	http.HandleFunc("/staff/delete", staffcontroller.Delete)

	http.HandleFunc("/transactions", transactioncontroller.Index)
	http.HandleFunc("/transactions/add", transactioncontroller.Add)
	http.HandleFunc("/transactions/edit", transactioncontroller.Edit)
	http.HandleFunc("/transactions/delete", transactioncontroller.Delete)

	log.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
