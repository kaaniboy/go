package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Problem struct {
	Question []byte
	Answer   []byte
}

var filename string = "data.txt"

func renderTemplate(w http.ResponseWriter, tmpl string, p *Problem) {
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, p)
}

func (p *Problem) save() error {
	return ioutil.WriteFile(filename, p.Question, 0600)
}

func loadProblem(title string) *Problem {
	body, _ := ioutil.ReadFile(filename)
	return &Problem{Question: body}
}
func handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", &Problem{Question: []byte("Test")})

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	question := []byte(r.FormValue("question"))
	p := &Problem{Question: question, Answer: []byte("")}
	p.save()
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
