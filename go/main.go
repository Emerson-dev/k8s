package main

import (
	"html/template"
	"net/http"
	"path"
)

type Challenge struct {
	Param  string
}

func main() {
	challenge := Challenge{"Code.education Rocks!"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Greeting(w, r, challenge)
	})
	http.ListenAndServe(":8000", nil)
}

func Greeting(w http.ResponseWriter, r *http.Request, c Challenge) {

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}