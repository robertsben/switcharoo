package main

import (
	"bytes"
	// "errors"
	// "encoding/xml"
	"html/template"
	"net/http"
	// "regexp"
)

type Page struct {
	Title string
	Body []byte
}

type Data struct {
	Conversion []byte
}

var templates = template.Must(template.ParseFiles("index.html", "output.html"))
// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")


func conversionHandler(w http.ResponseWriter, r *http.Request) {
	document := r.FormValue("document")
	NewDecoder(bytes.NewBufferString(document)).Decode(&Node{})
	data_out := &Data{Conversion: []byte(document)}
	renderTemplate(w, "output", data_out)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *Data) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", data) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", &Data{})
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/convert", conversionHandler)
	// http.HandleFunc("/edit/", makeHandler(editHandler))
	// http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}