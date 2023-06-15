package staffcontroller

import (
	"go-bookstore/models/staffmodel"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	staff := staffmodel.GetAll()
	data := map[string]any{
		"staff": staff,
	}

	temp, err := template.ParseFiles("views/staff/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
