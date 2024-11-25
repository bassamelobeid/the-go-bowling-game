package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", viewHandler)

	println("Local server running: http://localhost:4444")

	log.Fatal(http.ListenAndServe(":4444", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func loadPage(title string) (*Page, error) {
	return &Page{Title: title}, nil
}

type Page struct {
	Title string
	Body  []byte
}
