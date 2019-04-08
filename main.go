package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type NewsAggsPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggsPage{Title: "asdadasd", News: "dsds"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> dddd </h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":60", nil)
}
