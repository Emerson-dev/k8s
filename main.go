package main

import (
	"fmt"
	"net/http"
)

type Challenge struct {
	Param  string
}

func main() {
	challenge := Challenge{"Code.education Rocks!"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Greeting(w, r, challenge)
	})
	http.ListenAndServe(":8000",
		nil)
}

func Greeting(w http.ResponseWriter, r *http.Request, c Challenge) {
	fmt.Fprintf(w, "<b>%v</b>",c.Param)
}