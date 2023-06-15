package bookcontroller

import (
	"go-bookstore/entities"
	"go-bookstore/models/authormodel"
	"go-bookstore/models/bookmodel"
	"go-bookstore/models/genremodel"
	"go-bookstore/models/publishermodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	books := bookmodel.GetAll()

	data := map[string]any{
		"books": books,
	}

	temp, err := template.ParseFiles("views/books/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/books/create.html")
		if err != nil {
			panic(err)
		}

		genres := genremodel.GetAll()
		data := map[string]any{
			"genres": genres,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var author entities.Authors

		author.Name = r.FormValue("author_name")

		authormodel.Create(author)
	}

	if r.Method == "POST" {
		var publisher entities.Publishers

		publisher.Name = r.FormValue("publisher_name")

		publishermodel.Create(publisher)
	}

	if r.Method == "POST" {
		var book entities.Books

		genreId, err := strconv.Atoi(r.FormValue("genre_id"))
		if err != nil {
			panic(err)
		}

		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			panic(err)
		}

		stockQty, err := strconv.Atoi(r.FormValue("stock_qty"))
		if err != nil {
			panic(err)
		}

		book.Title = r.FormValue("title")
		book.Author.Name = r.FormValue("author_name")
		book.Publisher.Name = r.FormValue("publisher_name")
		book.Genre.Id = uint(genreId)
		book.PublicationDate = r.FormValue("publication_date")
		book.ISBN = r.FormValue("isbn")
		book.Price = float32(price)
		book.StockQty = int64(stockQty)
		book.CreatedAt = time.Now()
		book.UpdatedAt = time.Now()

		if ok := bookmodel.Create(book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	book := bookmodel.Detail(id)
	data := map[string]any{
		"book": book,
	}

	temp, err := template.ParseFiles("views/books/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/books/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		book := bookmodel.Detail(id)
		genres := genremodel.GetAll()

		data := map[string]any{
			"book":   book,
			"genres": genres,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var book entities.Books

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		genreId, err := strconv.Atoi(r.FormValue("genre_id"))
		if err != nil {
			panic(err)
		}

		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			panic(err)
		}

		stockQty, err := strconv.Atoi(r.FormValue("stock_qty"))
		if err != nil {
			panic(err)
		}

		book.Title = r.FormValue("title")
		book.Genre.Id = uint(genreId)
		book.Price = float32(price)
		book.StockQty = int64(stockQty)
		book.UpdatedAt = time.Now()

		if ok := bookmodel.Update(id, book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := bookmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
