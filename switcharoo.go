package main

import (
	// "errors"
	// "encoding/xml"
	"flag"
	"html/template"
	"io/ioutil"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	// "regexp"
)

type Page struct {
	Title string
	Body []byte
}

type Data struct {
	Conversion []byte
}

var Debug bool
var DumpFail bool

var templates = template.Must(template.ParseFiles("index.html", "output.html"))

func conversionHandler(w http.ResponseWriter, r *http.Request) {
	document := r.FormValue("document")
	output, err := Convert(document)

	if err != nil {
		renderTemplate(w, "output", &Data{Conversion: []byte(err.Error())})
	} else {
		data_out := &Data{Conversion: output.Bytes()}
		renderTemplate(w, "output", data_out)
	}
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

func fileConversionHandler(inputfilename string, outputfilename string) {
	input, _ := ioutil.ReadFile(inputfilename)
	output, err := Convert(string(input))

	if err != nil {
		if Debug {
			spew.Dump(err)
		}
		if DumpFail {
			ioutil.WriteFile("./failure.json", output.Bytes(), 0644)	
		}
		ioutil.WriteFile(outputfilename, []byte(err.Error()), 0644)
	} else {
		ioutil.WriteFile(outputfilename, output.Bytes(), 0644)
	}
}

func main() {
	var source string
	var destination string
	flag.StringVar(&source, "source", "", "path to the source xml file")
	flag.StringVar(&destination, "destination", "./example.json", "path to the output json file")
	flag.BoolVar(&Debug, "debug", false, "whether to add debugging logs")
	flag.BoolVar(&DumpFail, "dumpfail", true, "whether to write failed conversion to failure.json")
	flag.Parse()

	if len(source) < 1 {
		http.HandleFunc("/", indexHandler)
		http.HandleFunc("/convert", conversionHandler)
		http.ListenAndServe(":8080", nil)
	} else {
		fileConversionHandler(source, destination)
	}




}