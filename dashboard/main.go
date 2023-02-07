package main

import (
	"fmt"
	"github.com/rwirdemann/bffdashboard/bff"
	"html/template"
	"net/http"
)

type Result struct {
	Name   string
	Status int
}

type IndexData struct {
	Results []Result
}

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./dashboard/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", index)
	fmt.Printf("Dashboard started on http://localhost:8082")
	http.ListenAndServe(":8082", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	loginStatus, _ := bff.TestLogin()

	data := IndexData{}
	data.Results = append(data.Results, Result{"Marketplace Login", loginStatus})
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}