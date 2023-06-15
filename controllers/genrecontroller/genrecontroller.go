package genrecontroller

import (
	"go-bookstore/entities"
	"go-bookstore/models/genremodel"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	genre := genremodel.GetAll()
	data := map[string]any{
		"genre": genre,
	}

	temp, err := template.ParseFiles("views/genres/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/genres/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var genre entities.Genres

		genre.Name = r.FormValue("name")

		ok := genremodel.Create(genre)
		if !ok {
			temp, _ := template.ParseFiles("views/genres/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/genres", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/genres/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		genre := genremodel.Detail(id)
		data := map[string]any{
			"genres": genre,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var genre entities.Genres

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		genre.Name = r.FormValue("name")

		if ok := genremodel.Update(id, genre); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/genres", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := genremodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/genres", http.StatusSeeOther)
}
