package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Body []byte
}

var filename string = "data.txt"

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, p)
}

func (p *Page) save() error {
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	body, _ := ioutil.ReadFile(filename)
	return &Page{Body: body}
}
func handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", &Page{Body: []byte("Get Em All in a line")})

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
