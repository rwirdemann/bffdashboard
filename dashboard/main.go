package main

import (
	"fmt"
	"github.com/rwirdemann/bffdashboard/bff"
	"html/template"
	"net/http"
	"time"
)

type Result struct {
	Updated string
	Name    string
	Status  string
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
	var checks []bff.HealthCheck
	login := bff.LoginHealthCheck{}
	checks = append(checks, login)
	data := IndexData{}
	for _, check := range checks {
		result := check.Run()
		data.Results = append(data.Results, Result{Updated: time.Now().Format("2006-01-02 15:04:05"), Name: check.Description(), Status: result})
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}
