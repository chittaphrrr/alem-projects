package main

import (
	"html/template"
	"log"
	"net/http"
	"nsabyrov/ascii-art-web/internal"
)

type Error struct {
	IsError     bool
	ErrorStatus int
	ErrorText   string
}

type Data struct {
	Error  Error
	Result string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		err := render(w, Data{})
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		r.ParseForm()

		if r.FormValue("text") == "" || r.FormValue("banners") == "" {
			errorHandler(w, http.StatusBadRequest)
			return
		}
		result, status := internal.Convert(r.FormValue("text"), r.FormValue("banners"))
		if status != http.StatusOK {
			errorHandler(w, status)
			return
		}

		err := render(w, Data{Result: result})
		if err != nil {
			log.Println(err)
			errorHandler(w, http.StatusInternalServerError)

			return
		}
	default:
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	newData := Data{
		Error: Error{true, status, http.StatusText(status)},
	}

	if err := render(w, newData); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func render(w http.ResponseWriter, newData Data) error {
	tmpl, err := template.ParseFiles("./ui/templates/index.html")
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, newData); err != nil {
		return err
	}
	return nil
}
